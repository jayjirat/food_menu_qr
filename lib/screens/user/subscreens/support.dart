import 'package:flutter/material.dart';
import 'package:food_menu_qr/components/custom_divider.dart';
import 'package:food_menu_qr/components/user_stack.dart';

Widget support(BuildContext context) {
  return userStack(
      child: Padding(
    padding: const EdgeInsets.all(16),
    child: Column(
      children: [
        Text(
            "If you have any issues or questions regarding the FoodQR app, feel free to contact us via email or phone. Our team is available every day and will respond promptly to assist you with any concerns."),
        const SizedBox(
          height: 30,
        ),
        customDivider(context),
        const SizedBox(
          height: 25,
        ),
        buildSupportTap(
            context: context,
            title: "Send Us Your Message",
            subtitle: "We’re here to help",
            routeName: '/'),
        const SizedBox(
          height: 25,
        ),
        customDivider(context),
        const SizedBox(
          height: 25,
        ),
        buildSupportTap(
            context: context,
            title: "Help center",
            subtitle: "General Information",
            routeName: '/help-center'),
        const SizedBox(
          height: 25,
        ),
        customDivider(context),
      ],
    ),
  ));
}

Widget buildSupportTap(
    {required BuildContext context,
    required String title,
    required String subtitle,
    required String routeName}) {
  return InkWell(
    onTap: () {
      Navigator.pushNamed(context, routeName);
    },
    child: Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        RichText(
            text: TextSpan(
                text: '$title\n',
                style: TextStyle(
                    color: Theme.of(context).primaryColor,
                    fontSize: 20,
                    fontWeight: FontWeight.w600),
                children: [
              TextSpan(
                  text: subtitle,
                  style: TextStyle(fontSize: 14, fontWeight: FontWeight.normal))
            ])),
        Icon(
          Icons.arrow_forward_ios_outlined,
          color: Theme.of(context).colorScheme.secondary,
        )
      ],
    ),
  );
}
