import 'package:flutter/material.dart';

Widget ownerStack(
    {required BuildContext context,
    required Widget body,
    required String header,
    required String content,
    double? headerFontSize,
    Widget? logOutButton,
    Widget? floatingActionButton}) {
  final Widget actualLogOutButton = logOutButton ?? Container();
  return Scaffold(
      appBar: AppBar(
        toolbarHeight: 120,
        title: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                RichText(
                  text: TextSpan(
                      text: header,
                      style: TextStyle(
                          fontSize: headerFontSize ?? 28,
                          fontWeight: FontWeight.bold,
                          color: Colors.white),
                      children: [
                        TextSpan(
                            text: content,
                            style: TextStyle(
                                fontSize: 14,
                                color: Theme.of(context).colorScheme.secondary))
                      ]),
                ),
                actualLogOutButton
              ],
            )),
      ),
      body: body,
      floatingActionButton: floatingActionButton);
}
