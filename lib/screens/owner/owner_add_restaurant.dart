import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/components/label_input.dart';
import 'package:food_menu_qr/components/owner_stack.dart';
import 'package:food_menu_qr/components/user_stack.dart';

class OwnerAddRestaurant extends ConsumerStatefulWidget {
  const OwnerAddRestaurant({super.key});

  @override
  OwnerAddRestaurantState createState() => OwnerAddRestaurantState();
}

class OwnerAddRestaurantState extends ConsumerState<OwnerAddRestaurant> {
  final nameController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return ownerStack(
        header: "Add New restaurant\n",
        headerFontSize: 26,
        content: "Fill in the details to create\nyour restaurant",
        context: context,
        body: userStack(
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16),
            child: ListView(
              children: [
                Column(
                  children: [
                    inputWithLabel(
                        context: context,
                        controller: nameController,
                        label: "Restaurant Name",
                        hintText: "",
                        obscureText: false),
                    const SizedBox(
                      height: 20,
                    ),
                    ElevatedButton(
                        onPressed: () {},
                        style: ElevatedButton.styleFrom(
                            padding: const EdgeInsets.symmetric(
                                horizontal: 100, vertical: 10),
                            backgroundColor:
                                Theme.of(context).colorScheme.secondary),
                        child: Text(
                          "Create",
                          style: TextStyle(
                              color: Colors.white,
                              fontSize: 20,
                              fontWeight: FontWeight.bold),
                        ))
                  ],
                )
              ],
            ),
          ),
        ));
  }
}
