import 'dart:convert';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/models/user.dart';
import 'package:http/http.dart' as http;

class UserNotifier extends StateNotifier<User?> {
  UserNotifier() : super(null);
  final baseURL = "http://10.0.2.2:5678/api";

  Future<void> register(
      {required String email,
      required String username,
      required String password,
      required String dateOfBirth}) async {
    final body = {
      "email": email,
      "fullname": username,
      "password": password,
      "date_of_birth": dateOfBirth,
      "role": "user"
    };

    final header = {
      'Content-Type': 'application/json',
    };

    final response = await http.post(Uri.parse("$baseURL/register"),
        body: jsonEncode(body), headers: header);

    if (response.statusCode == 200) {
      print("Registration successful");
    } else {
      print(response.body);
      print("Registration failed: ${response.statusCode}");
      // เพิ่ม logic สำหรับการแสดง error message
    }
  }
}

final userNotifierProvider = StateNotifierProvider<UserNotifier, User?>((ref) {
  return UserNotifier();
});
