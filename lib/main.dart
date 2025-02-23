import 'package:flutter/material.dart';
import 'package:food_menu_qr/screens/homepage.dart';
import 'package:food_menu_qr/screens/login.dart';
import 'package:food_menu_qr/screens/register.dart';
import 'package:food_menu_qr/screens/user_screens/home.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData.light(
        useMaterial3: true,
      ).copyWith(
          primaryColor: Color(0xFF391713),
          colorScheme: ColorScheme.fromSeed(
            primary: const Color(0xFFF5CB58),
            seedColor: const Color(0xFFF5CB58),
            secondary: const Color(0xFFE95322),
          ),
          scaffoldBackgroundColor: Colors.white,
          elevatedButtonTheme: ElevatedButtonThemeData(
              style: ElevatedButton.styleFrom(
                  foregroundColor: Colors.black,
                  backgroundColor: Color(0xFFF5CB58),
                  padding:
                      const EdgeInsets.symmetric(horizontal: 100, vertical: 14),
                  textStyle:
                      TextStyle(fontSize: 16, fontWeight: FontWeight.w600))),
          appBarTheme: AppBarTheme(backgroundColor: Color(0xFFF5CB58))),
      initialRoute: '/',
      routes: {
        '/': (context) => const SplashScreen(),
        '/homepage': (context) => const HomePage(),
        '/login': (context) => const Login(),
        '/register': (context) => const Register(),
        //user screens
        '/home': (context) => const Home(),
      },
    );
  }
}

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  @override
  void initState() {
    super.initState();
    Future.delayed(const Duration(seconds: 2), () {
      if (mounted) {
        Navigator.pushReplacementNamed(context, '/homepage');
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).colorScheme.primary,
      body: Center(
        child: FadeTransition(
          opacity: AlwaysStoppedAnimation(1.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(
                Icons.fastfood,
                size: 200,
                color: Theme.of(context).colorScheme.secondary,
              ),
              const SizedBox(
                height: 20,
              ),
              RichText(
                  text: TextSpan(
                      text: "FOOD",
                      style: TextStyle(
                          color: Theme.of(context).colorScheme.secondary,
                          fontWeight: FontWeight.bold,
                          fontSize: 32),
                      children: [
                    TextSpan(text: "QR", style: TextStyle(color: Colors.white))
                  ])),
            ],
          ),
        ),
      ),
    );
  }
}
