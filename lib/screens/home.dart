import 'package:flutter/material.dart';

class Home extends StatefulWidget {
  const Home({super.key});

  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).colorScheme.secondary,
      body: Center(
        child: FadeTransition(
          opacity: AlwaysStoppedAnimation(1.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(
                Icons.fastfood,
                size: 200,
                color: Color(0xFFF5CB58),
              ),
              const SizedBox(
                height: 20,
              ),
              RichText(
                  text: TextSpan(
                      text: "FOOD",
                      style: TextStyle(
                          color: Color(0xFFF5CB58),
                          fontWeight: FontWeight.bold,
                          fontSize: 32),
                      children: [
                    TextSpan(text: "QR", style: TextStyle(color: Colors.white))
                  ])),
              const SizedBox(
                height: 30,
              ),
              Text(
                "FoodQR – Scan, Order, Enjoy!",
                style: TextStyle(fontSize: 14, color: Colors.white),
              ),
              const SizedBox(
                height: 40,
              ),
              ElevatedButton(
                onPressed: () {
                  Navigator.pushNamed(context, '/login');
                },
                child: Text("Sign In"),
              ),
              const SizedBox(
                height: 10,
              ),
              ElevatedButton(
                style: ElevatedButton.styleFrom(
                    backgroundColor: Color(0xFFF3E9B5)),
                onPressed: () {
                  Navigator.pushNamed(context, '/register');
                },
                child: Text("Sign Up"),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
