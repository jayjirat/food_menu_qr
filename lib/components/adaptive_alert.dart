import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

void adaptiveAlert(BuildContext ctx, {String title = '', String content = ''}) {
  if (Theme.of(ctx).platform == TargetPlatform.iOS) {
    showCupertinoDialog(
        context: ctx,
        builder: (ctx) => CupertinoAlertDialog(
              title: Text(title),
              content: Text(content),
              actions: [
                CupertinoDialogAction(
                  child: Text('Ok'),
                  onPressed: () => Navigator.of(ctx).pop(),
                ),
                CupertinoDialogAction(
                  onPressed: () => Navigator.of(ctx).pop(),
                  isDestructiveAction: true,
                  isDefaultAction: true,
                  child: Text('primary'),
                ),
              ],
            ));
  } else if (Theme.of(ctx).platform == TargetPlatform.android) {
    showDialog(
        context: ctx,
        builder: (ctx) => AlertDialog(
              title: Text(title),
              content: Text(content),
              actions: [
                TextButton(
                    onPressed: () {
                      Navigator.of(ctx).pop();
                    },
                    child: Text("ok"))
              ],
            ));
  }
}
