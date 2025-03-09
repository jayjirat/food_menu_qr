import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/screens/user/subscreens/history.dart';
import 'package:food_menu_qr/screens/user/subscreens/home.dart';
import 'package:food_menu_qr/screens/user/subscreens/support.dart';

class UserMain extends ConsumerStatefulWidget {
  const UserMain({super.key});

  @override
  UserMainState createState() => UserMainState();
}

class UserMainState extends ConsumerState<UserMain> {
  int selectedIndex = 0;

  @override
  Widget build(BuildContext context) {
    List<Widget> subScreens = [
      home(context),
      History(),
      home(context),
      home(context),
      support(context)
    ];
    return Scaffold(
      appBar: AppBar(
        toolbarHeight: 150,
        automaticallyImplyLeading: false,
        title: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              RichText(
                text: TextSpan(
                  text: "Welcome, ",
                  style: TextStyle(
                      fontSize: 26,
                      fontWeight: FontWeight.bold,
                      color: Colors.white),
                  children: [
                    TextSpan(
                      text: "Fullname\n",
                      style: TextStyle(
                          fontWeight: FontWeight.bold,
                          color: Theme.of(context).colorScheme.secondary),
                    ),
                    TextSpan(
                      text:
                          "Rise and shine! It's time to enjoy\n something delicious.",
                      style: TextStyle(
                          fontSize: 12,
                          color: Theme.of(context).colorScheme.secondary),
                    ),
                  ],
                ),
              ),
              const SizedBox(
                width: 60,
              ),
              InkWell(
                onTap: () {},
                child: Container(
                  height: 30,
                  width: 30,
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(20),
                  ),
                  child: Center(
                    child: Icon(
                      Icons.notifications_outlined,
                      color: Theme.of(context).colorScheme.secondary,
                      size: 25,
                    ),
                  ),
                ),
              ),
              const SizedBox(
                width: 10,
              ),
              InkWell(
                onTap: () {},
                child: Container(
                  height: 30,
                  width: 30,
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(20),
                  ),
                  child: Center(
                    child: Icon(
                      Icons.person_outline,
                      color: Theme.of(context).colorScheme.secondary,
                      size: 25,
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
      body: subScreens[selectedIndex],
      bottomNavigationBar: ClipRRect(
        borderRadius: BorderRadius.only(
            topLeft: Radius.circular(40), topRight: Radius.circular(40)),
        child: BottomNavigationBar(
          type: BottomNavigationBarType.fixed,
          backgroundColor: Theme.of(context).colorScheme.secondary,
          unselectedItemColor: Colors.white,
          selectedItemColor: Theme.of(context).colorScheme.primary,
          currentIndex: selectedIndex,
          onTap: (index) {
            setState(() {
              selectedIndex = index;
            });
          },
          items: [
            BottomNavigationBarItem(
              icon: Icon(Icons.home_outlined),
              label: "Home",
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.history_outlined),
              label: "History",
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.qr_code_outlined),
              label: "QR code",
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.favorite_outline),
              label: "Favorites",
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.support_agent_outlined),
              label: "Support",
            ),
          ],
        ),
      ),
    );
  }
}
