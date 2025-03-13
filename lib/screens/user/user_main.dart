import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/components/bottom_nav.dart';
import 'package:food_menu_qr/components/show_snackbar.dart';
import 'package:food_menu_qr/providers/user_provider.dart';
import 'package:food_menu_qr/screens/user/subscreens/history.dart';
import 'package:food_menu_qr/screens/user/subscreens/home.dart';
import 'package:food_menu_qr/screens/user/subscreens/qr_code.dart';
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
    final user = ref.read(userNotifierProvider)!;
    final scaffoldKey = GlobalKey<ScaffoldState>();
    List<Widget> subScreens = [
      home(context),
      History(),
      QrCode(),
      support(context)
    ];
    return Scaffold(
        key: scaffoldKey,
        appBar: AppBar(
          toolbarHeight: 150,
          automaticallyImplyLeading: false,
          title: Padding(
            padding: const EdgeInsets.all(8.0),
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Expanded(
                  child: RichText(
                    text: TextSpan(
                      text: "Welcome, ",
                      style: TextStyle(
                          fontSize: 26,
                          fontWeight: FontWeight.bold,
                          color: Colors.white),
                      children: [
                        TextSpan(
                          text: "${(user.username).toUpperCase()}\n",
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
                  onTap: () {
                    if (scaffoldKey.currentState!.isDrawerOpen) {
                      scaffoldKey.currentState!.closeDrawer();
                    } else {
                      scaffoldKey.currentState!.openDrawer();
                    }
                  },
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
        drawer: Drawer(
          shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.only(
                  topRight: Radius.circular(100),
                  bottomRight: Radius.circular(100))),
          backgroundColor: Theme.of(context).colorScheme.secondary,
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16),
            child: ListView(
              padding: EdgeInsets.zero,
              children: [
                SizedBox(
                  height: 150,
                  child: DrawerHeader(
                    padding: EdgeInsets.all(10),
                    child: ListTile(
                      contentPadding: EdgeInsets.zero,
                      title: RichText(
                        text: TextSpan(
                            text: "${(user.username).toUpperCase()}\n",
                            style: TextStyle(
                                fontSize: 28, fontWeight: FontWeight.bold),
                            children: [
                              TextSpan(
                                  // ignore: unnecessary_string_interpolations
                                  text: "${user.email}",
                                  style: TextStyle(
                                      fontSize: 16,
                                      fontWeight: FontWeight.normal,
                                      color: Color(0xFFF3E9B5)))
                            ]),
                      ),
                    ),
                  ),
                ),
                const SizedBox(
                  height: 14,
                ),
                drawerItem(
                    icon: Icons.person_outline,
                    label: "My Profile",
                    onTap: () {}),
                const SizedBox(
                  height: 14,
                ),
                drawerItem(
                    icon: Icons.settings_outlined,
                    label: "Settings",
                    onTap: () {}),
                const SizedBox(
                  height: 14,
                ),
                drawerItem(
                    icon: Icons.logout_outlined,
                    label: "Log Out",
                    onTap: () async {
                      final response = await ref
                          .read(userNotifierProvider.notifier)
                          .logout();
                      if (context.mounted) {
                        if (response["status"]) {
                          Navigator.pushNamedAndRemoveUntil(
                            context,
                            '/homepage',
                            (route) => false,
                          );
                        } else {
                          showSnackBar(context, response["message"]);
                        }
                      }
                    })
              ],
            ),
          ),
        ),
        bottomNavigationBar: bottomNav(
          context: context,
          selectedIndex: selectedIndex,
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
              icon: Icon(Icons.support_agent_outlined),
              label: "Support",
            ),
          ],
        ));
  }

  Widget drawerItem(
      {required IconData icon,
      required String label,
      required GestureTapCallback onTap}) {
    return ListTile(
        leading: Container(
          height: 40,
          width: 40,
          decoration: BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.circular(16),
          ),
          child: Center(
            child: Icon(
              icon,
              color: Theme.of(context).colorScheme.secondary,
              size: 30,
            ),
          ),
        ),
        title: Text(
          label,
          style: TextStyle(
              fontSize: 20,
              fontWeight: FontWeight.w500,
              color: Color(0xFFF3E9B5)),
        ),
        onTap: onTap);
  }
}
