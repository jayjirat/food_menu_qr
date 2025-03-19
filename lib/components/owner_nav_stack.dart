import 'package:flutter/material.dart';
import 'package:food_menu_qr/components/bottom_nav.dart';

Widget ownerNavStack(
    {required BuildContext context,
    required int selectedIndex,
    required Function(int) onTap,
    required Widget body}) {
  return Scaffold(
    appBar: AppBar(
      toolbarHeight: 80,
      title: Text("McDonald",
          style: TextStyle(
              fontSize: 26,
              fontWeight: FontWeight.bold,
              color: Theme.of(context).colorScheme.secondary)),
      centerTitle: true,
    ),
    body: body,
    bottomNavigationBar: bottomNav(
      context: context,
      selectedIndex: selectedIndex,
      onTap: onTap,
      items: [
        BottomNavigationBarItem(
          icon: Icon(Icons.dashboard_outlined),
          label: "Dashboard",
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.restaurant_menu_outlined),
          label: "Menu",
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.list_alt_outlined),
          label: "Orders",
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.store_outlined),
          label: "Manage",
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.support_agent_outlined),
          label: "Support",
        ),
      ],
    ),
  );
}
