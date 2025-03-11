import 'dart:convert';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/models/user.dart';
import 'package:http/http.dart' as http;

class UserNotifier extends StateNotifier<User?> {
  UserNotifier() : super(null);
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
      final bodyResonse = jsonDecode(response.body) as Map<String, dynamic>;
      if (response.statusCode == 200) {
        token = response.headers["authorization"]!.split(" ")[1];
        Map<String, dynamic> user = bodyResonse["user"];
        state = User(
            username: user["fullname"],
            email: user["email"],
            dateOfBirth: user["date_of_birth"],
            password: user["password"],
            role: user["role"]);
        return {"status": true, "message": "Login successful"};
      } else {
        return {"status": false, "message": bodyResonse["message"]};
      }
    } catch (e) {
      return {"status": false, "message": e};
    }
  }
}

final userNotifierProvider = StateNotifierProvider<UserNotifier, User?>((ref) {
  return UserNotifier();
});
