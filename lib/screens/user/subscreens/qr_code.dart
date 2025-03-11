import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:food_menu_qr/components/user_stack.dart';
import 'package:mobile_scanner/mobile_scanner.dart';
import 'package:permission_handler/permission_handler.dart';

class QrCode extends ConsumerStatefulWidget {
  const QrCode({super.key});

  @override
  QrCodeState createState() => QrCodeState();
}

class QrCodeState extends ConsumerState<QrCode> {
  final MobileScannerController scannerController = MobileScannerController();

  Future<void> requestCameraPermission() async {
    if (await Permission.camera.request().isGranted) {
      debugPrint("✅ Camera permission granted");
    } else {
      debugPrint("❌ Camera permission denied");
    }
  }

  void onDetect(BarcodeCapture capture) {
    final barcode = capture.barcodes.first;
    if (barcode.format == BarcodeFormat.qrCode) {
      final String? value = barcode.rawValue;
      if (value != null) {
        debugPrint("QR Code Detected: $value");
      }
    } else {
      debugPrint("This is not a QR Code.");
    }
  }

  @override
  void initState() {
    super.initState();
    requestCameraPermission();
  }

  @override
  Widget build(BuildContext context) {
    return userStack(
      child: Stack(
        children: [
          MobileScanner(
            controller: scannerController,
            onDetect: onDetect,
          ),
          Center(
              child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Container(
                width: 250,
                height: 250,
                decoration: BoxDecoration(
                  border: Border.all(color: Colors.black45, width: 4),
                  borderRadius: BorderRadius.circular(16),
                ),
              ),
              const SizedBox(
                height: 40,
              ),
              ElevatedButton(
                onPressed: () => scannerController.switchCamera(),
                child: const Text('Switch Camera'),
              ),
            ],
          )),
        ],
      ),
    );
  }
}
