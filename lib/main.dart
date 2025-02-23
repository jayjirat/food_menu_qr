import 'package:flutter/material.dart';
import 'package:food_menu_qr/screens/home.dart';
import 'package:food_menu_qr/screens/login.dart';
import 'package:food_menu_qr/screens/register.dart';

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
              seedColor: const Color(0xFF391713),
              secondary: const Color(0xFFE95322)),
          scaffoldBackgroundColor: Colors.white,
          elevatedButtonTheme: ElevatedButtonThemeData(
              style: ElevatedButton.styleFrom(
                  foregroundColor: Colors.black,
                  backgroundColor: Color(0xFFFFF5E1),
                  padding:
                      const EdgeInsets.symmetric(horizontal: 100, vertical: 14),
                  textStyle:
                      TextStyle(fontSize: 16, fontWeight: FontWeight.w600))),
          appBarTheme: AppBarTheme(backgroundColor: Color(0xFFF5CB58))),
      initialRoute: '/',
      routes: {
        '/': (context) => const SplashScreen(),
        '/home': (context) => const Home(),
        '/login': (context) => const Login(),
        '/register': (context) => const Register()
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
    Future.delayed(const Duration(seconds: 3), () {
      if (mounted) {
        Navigator.pushReplacementNamed(context, '/home');
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: FadeTransition(
          opacity: AlwaysStoppedAnimation(1.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              ClipRRect(
                borderRadius: BorderRadius.circular(50),
                child: Image.network(
                  "https://fastly.picsum.photos/id/362/200/200.jpg?hmac=AKqfQ8tnyGapdUtZ1f35ugad3WkJY-g1tn5hi7kF2zY",
                  width: 150,
                ),
              ),
              const SizedBox(
                height: 20,
              ),
              Text(
                "FoodQR",
                style: TextStyle(fontSize: 32, fontWeight: FontWeight.bold),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
