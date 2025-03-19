import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/components/custom_divider.dart';
import 'package:food_menu_qr/components/owner_stack.dart';
import 'package:food_menu_qr/components/show_snackbar.dart';
import 'package:food_menu_qr/components/user_stack.dart';
import 'package:food_menu_qr/providers/user_provider.dart';

class OwnerMain extends ConsumerStatefulWidget {
  const OwnerMain({super.key});

  @override
  OwnerMainState createState() => OwnerMainState();
}

class OwnerMainState extends ConsumerState<OwnerMain> {
  int selectedIndex = 0;

  @override
  Widget build(BuildContext context) {
    final user = ref.read(userNotifierProvider)!;
    return ownerStack(
      header: "Hi ${(user.username).toUpperCase()}\n",
      content:
          "Quickly manage your restaurants\nand provide a better customer experience!",
      logOutButton: InkWell(
        onTap: () async {
          final response =
              await ref.read(userNotifierProvider.notifier).logout();
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
        },
        child: Container(
          height: 40,
          width: 40,
          decoration: BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.circular(20),
          ),
          child: Center(
            child: Icon(
              Icons.logout_outlined,
              color: Theme.of(context).colorScheme.secondary,
              size: 25,
            ),
          ),
        ),
      ),
      context: context,
      body: userStack(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16),
          child: ListView(
            children: [
              Column(
                children: [
                  restaurantList(context: context),
                  restaurantList(context: context),
                  restaurantList(context: context),
                  restaurantList(context: context),
                  restaurantList(context: context),
                ],
              )
            ],
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => Navigator.pushNamed(context, '/owner-add-restaurant'),
        backgroundColor: Theme.of(context).colorScheme.primary,
        child: Icon(
          Icons.add_business_outlined,
          size: 40,
          color: Colors.white,
        ),
      ),
    );
  }

  Widget restaurantList({required BuildContext context}) {
    return Column(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(
              "McDonald",
              style: TextStyle(
                  color: Theme.of(context).primaryColor,
                  fontSize: 20,
                  fontWeight: FontWeight.w700),
            ),
            ElevatedButton(
                onPressed: () {},
                style: ElevatedButton.styleFrom(
                    padding: const EdgeInsets.all(0),
                    backgroundColor: Theme.of(context).colorScheme.secondary),
                child: Icon(
                  Icons.arrow_forward_ios_outlined,
                  color: Colors.white,
                )),
          ],
        ),
        customDivider(context),
        const SizedBox(
          height: 10,
        )
      ],
    );
  }
}


// ownerNavStack(
//         context: context,
//         selectedIndex: selectedIndex,
//         onTap: (index) {
//           setState(() {
//             selectedIndex = index;
//           });
//         },
//         body: userStack(
//             child: Padding(
//           padding: const EdgeInsets.all(16),
//           child: ListView(
//             children: [
//               Column(
//                 children: [],
//               )
//             ],
//           ),
//         )));