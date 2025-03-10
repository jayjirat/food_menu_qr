import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/components/adaptive_alert.dart';
import 'package:food_menu_qr/components/label_input.dart';
import 'package:food_menu_qr/components/main_stack.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:food_menu_qr/components/show_snackbar.dart';
import 'package:food_menu_qr/providers/user_provider.dart';

class Register extends ConsumerStatefulWidget {
  const Register({super.key});

  @override
  RegisterState createState() => RegisterState();
}

class RegisterState extends ConsumerState<Register> {
  final formKey = GlobalKey<FormState>();
  final fullnameController = TextEditingController();
  final passwordController = TextEditingController();
  final confirmPasswordController = TextEditingController();
  final emailController = TextEditingController();
  final dateOfBirthController = TextEditingController();

  final String termsOfUse = """
By accessing or using this app, you agree to be bound by these Terms of Use. If you do not agree to the terms, please do not use the app. We reserve the right to modify or update these terms at any time, so please review them periodically.

1. Usage Rights: You are granted a limited, non-transferable license to use the app for personal, non-commercial purposes.
2. Prohibited Activities: You may not use the app for illegal or unauthorized activities.
3. Privacy: Your use of the app is subject to our Privacy Policy, which can be found separately.
4. Limitation of Liability: We are not responsible for any damages that may arise from using the app.
""";

  final String privacyPolicy = """
We value your privacy. This policy outlines how we collect, use, and protect your personal information.

1. Information Collection: We collect personal information such as your name and email address when you register for the app.
2. Information Use: Your information is used to personalize your experience and to provide services.
3. Data Security: We implement security measures to protect your personal information.
4. Third-Party Sharing: We do not sell or share your information with third parties, except as required by law.
5. Your Rights: You can update or delete your personal information at any time.
""";

  @override
  Widget build(BuildContext context) {
    return mainStack(
        context: context,
        title: "Sign up",
        child: ListView(children: [
          Padding(
              padding: const EdgeInsets.symmetric(horizontal: 24),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Form(
                      key: formKey,
                      child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            inputWithLabel(
                              context: context,
                              controller: fullnameController,
                              label: "Full name",
                              hintText: "John Doe",
                              obscureText: false,
                            ),
                            const SizedBox(
                              height: 10,
                            ),
                            inputWithLabel(
                              context: context,
                              controller: passwordController,
                              label: "Password",
                              hintText: "*************",
                              obscureText: true,
                            ),
                            const SizedBox(
                              height: 10,
                            ),
                            inputWithLabel(
                              context: context,
                              controller: confirmPasswordController,
                              label: "Confirm Password",
                              hintText: "*************",
                              obscureText: true,
                            ),
                            const SizedBox(
                              height: 10,
                            ),
                            inputWithLabel(
                                context: context,
                                controller: emailController,
                                label: "Email",
                                hintText: "example@example.com",
                                obscureText: false,
                                textInputType: TextInputType.emailAddress),
                            const SizedBox(
                              height: 10,
                            ),
                            inputWithLabel(
                                context: context,
                                controller: dateOfBirthController,
                                label: "Date of birth",
                                hintText: "DD / MM /YYY",
                                obscureText: false,
                                textInputType: TextInputType.datetime),
                            const SizedBox(
                              height: 20,
                            ),
                            Center(
                              child: RichText(
                                  textAlign: TextAlign.center,
                                  text: TextSpan(
                                      text: "By continuing, you agree to\n",
                                      style: TextStyle(
                                          color: Colors.black, fontSize: 12),
                                      children: [
                                        TextSpan(
                                            text: "Terms of Use ",
                                            style: TextStyle(
                                                fontWeight: FontWeight.bold,
                                                color: Theme.of(context)
                                                    .colorScheme
                                                    .secondary),
                                            recognizer: TapGestureRecognizer()
                                              ..onTap = () => adaptiveAlert(
                                                  context,
                                                  title: "Terms of Use",
                                                  content: termsOfUse)),
                                        TextSpan(text: "and "),
                                        TextSpan(
                                            text: "Privacy Policy.",
                                            style: TextStyle(
                                                fontWeight: FontWeight.bold,
                                                color: Theme.of(context)
                                                    .colorScheme
                                                    .secondary),
                                            recognizer: TapGestureRecognizer()
                                              ..onTap = () => adaptiveAlert(
                                                  context,
                                                  title: "Privacy Policy",
                                                  content: privacyPolicy))
                                      ])),
                            ),
                            const SizedBox(
                              height: 10,
                            ),
                            Center(
                              child: ElevatedButton(
                                  style: ElevatedButton.styleFrom(
                                      backgroundColor: Theme.of(context)
                                          .colorScheme
                                          .secondary),
                                  onPressed: () async {
                                    if (formKey.currentState?.validate() ??
                                        false) {
                                      formKey.currentState?.save();
                                      handleRegister(context);
                                    }
                                  },
                                  child: Text(
                                    "Sign Up",
                                    style: TextStyle(
                                        color: Colors.white, fontSize: 20),
                                  )),
                            ),
                            const SizedBox(
                              height: 10,
                            ),
                            Center(
                              child: Column(
                                children: [
                                  Text("or sign up with"),
                                  Row(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                      iconButton(FontAwesomeIcons.google),
                                      iconButton(FontAwesomeIcons.facebook),
                                      iconButton(Icons.face_6),
                                    ],
                                  ),
                                  Center(
                                    child: RichText(
                                        text: TextSpan(
                                            text: "Already have an account? ",
                                            style:
                                                TextStyle(color: Colors.black),
                                            children: [
                                          TextSpan(
                                              text: "Sign In",
                                              style: TextStyle(
                                                  color: Theme.of(context)
                                                      .colorScheme
                                                      .secondary),
                                              recognizer: TapGestureRecognizer()
                                                ..onTap = () {
                                                  Navigator.pushNamed(
                                                      context, '/login');
                                                }),
                                        ])),
                                  )
                                ],
                              ),
                            )
                          ]))
                ],
              ))
        ]));
  }

  IconButton iconButton(IconData icon) {
    return IconButton(
      color: Theme.of(context).colorScheme.secondary,
      onPressed: () {},
      icon: Icon(icon),
    );
  }

  void handleRegister(BuildContext context) async {
    if (passwordController.text == confirmPasswordController.text) {
      final response = await ref.read(userNotifierProvider.notifier).register(
          email: emailController.text,
          username: fullnameController.text,
          password: passwordController.text,
          dateOfBirth: dateOfBirthController.text);
      if (context.mounted) {
        if (response["status"]) {
          Navigator.pushReplacementNamed(context, '/login',
              arguments: {"message": response["message"]});
        } else {
          showSnackBar(context, response["message"]);
        }
      }
    } else {
      showSnackBar(context, "Password does not match");
      passwordController.clear();
      confirmPasswordController.clear();
    }
  }
}
