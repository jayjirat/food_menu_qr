import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:food_menu_qr/components/main_stack.dart';

class HelpCenter extends ConsumerStatefulWidget {
  const HelpCenter({super.key});

  @override
  HelpCenterState createState() => HelpCenterState();
}

class HelpCenterState extends ConsumerState<HelpCenter> {
  int selectedIndex = 1;
  int faqSelectedIndex = 1;
  @override
  Widget build(BuildContext context) {
    return mainStack(
        context: context,
        title: "Help Center",
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 30),
          child: Column(
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  topButton(
                      context: context,
                      text: "FAQ",
                      width: 165,
                      onPressed: () => setState(() {
                            selectedIndex = 1;
                          }),
                      style: selectedIndex == 1
                          ? selectedStyle()
                          : unselectedStyle()),
                  topButton(
                      context: context,
                      text: "Contact Us",
                      width: 165,
                      onPressed: () => setState(() {
                            selectedIndex = 2;
                          }),
                      style: selectedIndex == 2
                          ? selectedStyle()
                          : unselectedStyle())
                ],
              ),
              selectedIndex == 1
                  ? const SizedBox(
                      height: 15,
                    )
                  : const SizedBox(
                      height: 35,
                    ),
              selectedIndex == 1 ? faq(context) : contactUs()
            ],
          ),
        ));
  }

  Widget faq(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        topButton(
            context: context,
            text: "General",
            width: 110,
            onPressed: () => setState(() {
                  faqSelectedIndex = 1;
                }),
            style: faqSelectedIndex == 1 ? selectedStyle() : unselectedStyle()),
        topButton(
            context: context,
            text: "Account",
            width: 110,
            onPressed: () => setState(() {
                  faqSelectedIndex = 2;
                }),
            style: faqSelectedIndex == 2 ? selectedStyle() : unselectedStyle()),
        topButton(
            context: context,
            text: "Services",
            width: 110,
            onPressed: () => setState(() {
                  faqSelectedIndex = 3;
                }),
            style: faqSelectedIndex == 3 ? selectedStyle() : unselectedStyle())
      ],
    );
  }

  Widget contactUs() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        contactUsElem(
            icon: Icons.headphones_outlined, text: "Customer service"),
        const SizedBox(
          height: 20,
        ),
        contactUsElem(icon: Icons.language_outlined, text: "Website"),
        const SizedBox(
          height: 20,
        ),
        contactUsElem(icon: Icons.facebook_outlined, text: "Facebook"),
        const SizedBox(
          height: 20,
        ),
        contactUsElem(icon: FontAwesomeIcons.instagram, text: "Instragram"),
      ],
    );
  }

  Widget contactUsElem({required IconData icon, required String text}) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Row(
          children: [
            Icon(
              icon,
              color: Theme.of(context).colorScheme.secondary,
              size: 40,
            ),
            const SizedBox(
              width: 20,
            ),
            Text(
              text,
              style: TextStyle(
                  color: Theme.of(context).primaryColor,
                  fontSize: 20,
                  fontWeight: FontWeight.w500),
            )
          ],
        ),
        Icon(Icons.arrow_drop_down_outlined),
      ],
    );
  }

  Widget topButton(
      {required BuildContext context,
      required String text,
      required double width,
      required VoidCallback onPressed,
      required ButtonStyle style}) {
    return SizedBox(
      height: 30,
      width: width,
      child:
          ElevatedButton(onPressed: onPressed, style: style, child: Text(text)),
    );
  }

  ButtonStyle selectedStyle() {
    return ElevatedButton.styleFrom(
        padding: const EdgeInsets.all(0),
        foregroundColor: Colors.white,
        backgroundColor: Theme.of(context).colorScheme.secondary);
  }

  ButtonStyle unselectedStyle() {
    return ElevatedButton.styleFrom(
        padding: const EdgeInsets.all(0),
        foregroundColor: Theme.of(context).colorScheme.secondary,
        backgroundColor: Color(0xFFFFDECF));
  }
}
