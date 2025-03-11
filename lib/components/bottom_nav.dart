import 'package:flutter/material.dart';

Widget bottomNav(
    {required BuildContext context,
    required int selectedIndex,
    required Function(int) onTap,
    required List<BottomNavigationBarItem> items}) {
  return ClipRRect(
    borderRadius: BorderRadius.only(
        topLeft: Radius.circular(40), topRight: Radius.circular(40)),
    child: BottomNavigationBar(
      type: BottomNavigationBarType.fixed,
      backgroundColor: Theme.of(context).colorScheme.secondary,
      unselectedItemColor: Colors.white,
      selectedItemColor: Theme.of(context).colorScheme.primary,
      currentIndex: selectedIndex,
      onTap: onTap,
      items: items,
    ),
  );
}
