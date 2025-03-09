import 'package:flutter/material.dart';

Widget filterButton(
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
