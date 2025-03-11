import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/components/owner_stack.dart';
import 'package:food_menu_qr/components/user_stack.dart';

class OwnerMain extends ConsumerStatefulWidget {
  const OwnerMain({super.key});

  @override
  OwnerMainState createState() => OwnerMainState();
}

class OwnerMainState extends ConsumerState<OwnerMain> {
  int selectedIndex = 1;
  @override
  Widget build(BuildContext context) {
    return ownerStack(
        context: context,
        selectedIndex: selectedIndex,
        onTap: (index) {
          setState(() {
            selectedIndex = index;
          });
        },
        body: userStack(child: Container()));
  }
}
