import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/components/label_input.dart';
import 'package:food_menu_qr/components/main_stack.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class Login extends ConsumerStatefulWidget {
  const Login({super.key});

  @override
  LoginState createState() => LoginState();
}

class LoginState extends ConsumerState<Login> {
  final formKey = GlobalKey<FormState>();
  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  String email = '';
  String password = '';
  @override
  Widget build(BuildContext context) {
    return mainStack(
      context: context,
      title: "Sign In",
      child: ListView(physics: NeverScrollableScrollPhysics(), children: [
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 24),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                "Welcome",
                style: TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                    color: Theme.of(context).primaryColor),
              ),
              const SizedBox(
                height: 10,
              ),
              Text(
                "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
                style: TextStyle(
                  fontSize: 13,
                ),
              ),
              const SizedBox(
                height: 40,
              ),
              Form(
                  key: formKey,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      inputWithLabel(
                          onSaved: (value) => email = value ?? '',
                          obscureText: false,
                          context: context,
                          controller: emailController,
                          label: "Email",
                          hintText: "example@example.com"),
                      const SizedBox(
                        height: 20,
                      ),
                      inputWithLabel(
                          onSaved: (value) => password = value ?? '',
                          obscureText: true,
                          context: context,
                          controller: passwordController,
                          label: "Password",
                          hintText: "*************"),
                      const SizedBox(
                        height: 20,
                      ),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.end,
                        children: [
                          Text(
                            "Forget Password",
                            style: TextStyle(
                              color: Theme.of(context).colorScheme.secondary,
                              fontWeight: FontWeight.w500,
                            ),
                          ),
                        ],
                      ),
                      const SizedBox(
                        height: 60,
                      ),
                      Center(
                        child: ElevatedButton(
                            style: ElevatedButton.styleFrom(
                                backgroundColor:
                                    Theme.of(context).colorScheme.secondary),
                            onPressed: () {
                              if (formKey.currentState?.validate() ?? false) {
                                formKey.currentState?.save();
                                ScaffoldMessenger.of(context) //Mock
                                    .showSnackBar(SnackBar(
                                        content: Row(
                                  children: [
                                    Text("Email: ${emailController.text}"),
                                    const SizedBox(
                                      width: 10,
                                    ),
                                    Text(
                                        "Password:  ${passwordController.text}")
                                  ],
                                )));
                              }
                            },
                            child: Text(
                              "Sign In",
                              style:
                                  TextStyle(color: Colors.white, fontSize: 20),
                            )),
                      )
                    ],
                  )),
              const SizedBox(
                height: 30,
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
                    const SizedBox(
                      height: 20,
                    ),
                    Center(
                      child: RichText(
                          text: TextSpan(
                              text: "Don’t have an account? ",
                              style: TextStyle(color: Colors.black),
                              children: [
                            TextSpan(
                                text: "Sign up",
                                style: TextStyle(
                                    color: Theme.of(context)
                                        .colorScheme
                                        .secondary),
                                recognizer: TapGestureRecognizer()
                                  ..onTap = () {
                                    Navigator.pushNamed(context, '/register');
                                  }),
                          ])),
                    )
                  ],
                ),
              )
            ],
          ),
        ),
      ]),
    );
  }

  IconButton iconButton(IconData icon) {
    return IconButton(
      color: Theme.of(context).colorScheme.secondary,
      onPressed: () {},
      icon: Icon(icon),
    );
  }
}
