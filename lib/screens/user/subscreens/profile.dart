// ignore_for_file: unnecessary_string_interpolations

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/components/filter_button.dart';
import 'package:food_menu_qr/components/main_stack.dart';
import 'package:food_menu_qr/providers/user_provider.dart';
import 'package:bcrypt/bcrypt.dart';

class Profile extends ConsumerStatefulWidget {
  const Profile({super.key});

  @override
  ProfileState createState() => ProfileState();
}

class ProfileState extends ConsumerState<Profile> {
  late TextEditingController oldPasswordController;
  @override
  void initState() {
    super.initState();
    oldPasswordController = TextEditingController();
    oldPasswordController.addListener(() {
      ref.read(passwordProvider.notifier).state = oldPasswordController.text;
    });
  }

  bool changePassword = false;
  bool isMatched = false;
  @override
  Widget build(BuildContext context) {
    final user = ref.read(userNotifierProvider)!;
    final emailController = TextEditingController(text: "${user.email}");
    final usernameController = TextEditingController(text: "${user.username}");
    final dateOfBirthController =
        TextEditingController(text: "${user.dateOfBirth}");
    final oldPassword = ref.watch(passwordProvider);
    final newPasswordController = TextEditingController();
    final newPasswordConfirmController = TextEditingController();
    if (BCrypt.checkpw(oldPassword, user.password)) {
      setState(() {
        isMatched = true;
      });
    }

    return mainStack(
        context: context,
        title: "My Profile",
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 20),
          child: ListView(children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                myInputProfile(
                    context: context,
                    controller: emailController,
                    label: "Email",
                    enabled: false,
                    obscureText: false),
                const SizedBox(
                  height: 20,
                ),
                myInputProfile(
                    context: context,
                    controller: usernameController,
                    label: "Full Name",
                    enabled: true,
                    obscureText: false),
                const SizedBox(
                  height: 20,
                ),
                myInputProfile(
                    context: context,
                    controller: dateOfBirthController,
                    label: "Date Of Birth",
                    enabled: true,
                    obscureText: false,
                    textInputType: TextInputType.datetime),
                if (changePassword && !isMatched)
                  Column(
                    children: [
                      const SizedBox(
                        height: 20,
                      ),
                      myInputProfile(
                        context: context,
                        controller: oldPasswordController,
                        label: "Enter Current Password",
                        enabled: true,
                        obscureText: true,
                      )
                    ],
                  ),
                if (isMatched && changePassword)
                  Column(
                    children: [
                      const SizedBox(
                        height: 20,
                      ),
                      myInputProfile(
                        context: context,
                        controller: newPasswordController,
                        label: "New Password",
                        enabled: true,
                        obscureText: true,
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                      myInputProfile(
                        context: context,
                        controller: newPasswordConfirmController,
                        label: "Confirm New Password",
                        enabled: true,
                        obscureText: true,
                      )
                    ],
                  ),
                const SizedBox(
                  height: 40,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceAround,
                  children: [
                    filterButton(
                        context: context,
                        text: !changePassword ? "Change Password" : "Cancel",
                        width: 150,
                        onPressed: () {
                          setState(() {
                            changePassword = !changePassword;
                            if (isMatched) {
                              oldPasswordController.clear();
                              isMatched = false;
                            }
                          });
                        },
                        style: changePassword
                            ? selectedStyle()
                            : unselectedStyle()),
                    filterButton(
                        context: context,
                        text: "Update Profile",
                        width: 150,
                        onPressed: () {},
                        style: !changePassword
                            ? selectedStyle()
                            : unselectedStyle())
                  ],
                )
              ],
            ),
          ]),
        ));
  }

  ButtonStyle selectedStyle() {
    return ElevatedButton.styleFrom(
        padding: const EdgeInsets.all(0),
        foregroundColor: Colors.white,
        backgroundColor: Theme.of(context).colorScheme.secondary);
  }

  ButtonStyle unselectedStyle() {
    return ElevatedButton.styleFrom(
        padding: const EdgeInsets.all(0),
        foregroundColor: Theme.of(context).colorScheme.secondary,
        backgroundColor: Color(0xFFFFDECF));
  }

  Widget myInputProfile({
    required BuildContext context,
    required TextEditingController controller,
    required String label,
    String? hintText,
    required bool enabled,
    required bool obscureText,
    TextInputType? textInputType,
  }) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          label,
          style: TextStyle(
            fontSize: 18,
            color: Theme.of(context).primaryColor,
            fontWeight: FontWeight.w600,
          ),
        ),
        const SizedBox(height: 5),
        TextField(
          enabled: enabled,
          controller: controller,
          obscureText: obscureText,
          keyboardType: textInputType ?? TextInputType.text,
          style: TextStyle(
            color: Theme.of(context).primaryColor,
            fontWeight: FontWeight.bold,
          ),
          decoration: InputDecoration(
            hintText: hintText,
            border: OutlineInputBorder(
              borderSide: BorderSide.none,
              borderRadius: BorderRadius.circular(16),
            ),
            filled: true,
            fillColor: const Color(0xFFF3E9B5),
            contentPadding: const EdgeInsets.symmetric(
              horizontal: 16,
              vertical: 12,
            ),
          ),
        ),
      ],
    );
  }
}
