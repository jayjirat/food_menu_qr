import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class OwnerMain extends ConsumerStatefulWidget {
  const OwnerMain({super.key});

  @override
  OwnerMainState createState() => OwnerMainState();
}

class OwnerMainState extends ConsumerState<OwnerMain> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(child: Text('OwnerMain')),
    );
  }
}
