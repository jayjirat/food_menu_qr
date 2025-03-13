import 'dart:convert';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/models/user.dart';
import 'package:http/http.dart' as http;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class UserNotifier extends StateNotifier<User?> {
  UserNotifier() : super(null);
  final storage = FlutterSecureStorage();
  final baseURL = "http://10.0.2.2:5678/api";
  String token = "";

  Future<Map<String, dynamic>> register({
    required String email,
    required String username,
    required String password,
    required String dateOfBirth,
  }) async {
    final body = {
      "email": email,
      "fullname": username,
      "password": password,
      "date_of_birth": dateOfBirth,
      "role": "user",
    };

    final header = {
      'Content-Type': 'application/json',
    };

    try {
      final response = await http.post(
        Uri.parse("$baseURL/register"),
        body: jsonEncode(body),
        headers: header,
      );

      if (response.statusCode == 200) {
        return {
          "status": true,
          "message":
              jsonDecode(response.body)["message"] ?? "Registration successful",
        };
      } else {
        return {
          "status": false,
          "message":
              jsonDecode(response.body)["message"] ?? "Registration failed",
        };
      }
    } catch (e) {
      return {
        "status": false,
        "message": "An error occurred: $e",
      };
    }
  }

  Future<Map<String, dynamic>> login({
    required String email,
    required String password,
  }) async {
    final body = {
      "email": email,
      "password": password,
    };

    final header = {
      'Content-Type': 'application/json',
    };

    try {
      final response = await http.post(
        Uri.parse("$baseURL/login"),
        body: jsonEncode(body),
        headers: header,
      );
      final bodyResponse = jsonDecode(response.body) as Map<String, dynamic>;
      if (response.statusCode == 200) {
        token = response.headers["authorization"]!.split(" ")[1];
        await storage.write(key: "token", value: token);
        final user = bodyResponse["user"];
        state = User(
            username: user["fullname"],
            email: user["email"],
            dateOfBirth: user["date_of_birth"],
            password: user["password"],
            role: user["role"]);
        return {"status": true, "message": "Login successful"};
      } else {
        return {"status": false, "message": bodyResponse["message"]};
      }
    } catch (e) {
      return {"status": false, "message": e};
    }
  }

  Future<Map<String, dynamic>> logout() async {
    try {
      await storage.delete(key: "token");
      state = null;
      return {
        "status": true,
        "message": "Logout successful",
      };
    } catch (e) {
      return {
        "status": false,
        "message": "Error while logging out. Please try again",
      };
    }
  }
}

final userNotifierProvider = StateNotifierProvider<UserNotifier, User?>((ref) {
  return UserNotifier();
});
