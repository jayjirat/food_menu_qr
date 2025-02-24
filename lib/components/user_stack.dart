import 'package:flutter/material.dart';

Widget userStack({required Widget child}) {
  return Center(
      child: Stack(children: [
    Container(
      width: double.infinity,
      height: double.infinity,
      color: Color(0xFFF5CB58),
    ),
    Container(
        width: double.infinity,
        height: double.infinity,
        padding: const EdgeInsets.only(top: 30),
        decoration: BoxDecoration(
            borderRadius: const BorderRadius.only(
                topLeft: Radius.circular(30), topRight: Radius.circular(30)),
            color: Colors.white),
        child: Padding(padding: const EdgeInsets.all(16), child: child))
  ]));
}
