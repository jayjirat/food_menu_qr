import 'package:flutter/material.dart';
import 'package:food_menu_qr/components/custom_divider.dart';
import 'package:food_menu_qr/components/filter_button.dart';
import 'package:food_menu_qr/components/user_stack.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class History extends ConsumerStatefulWidget {
  const History({super.key});

  @override
  HistoryState createState() => HistoryState();
}

class HistoryState extends ConsumerState<History> {
  int selectedIndex = 1;
  @override
  Widget build(BuildContext context) {
    return userStack(
        child: ListView(children: [
      Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: [
          filterButton(
              context: context,
              text: "Active",
              width: 110,
              onPressed: () => setState(() {
                    selectedIndex = 1;
                  }),
              style: selectedIndex == 1 ? selectedStyle() : unselectedStyle()),
          filterButton(
              context: context,
              text: "Completed",
              width: 110,
              onPressed: () => setState(() {
                    selectedIndex = 2;
                  }),
              style: selectedIndex == 2 ? selectedStyle() : unselectedStyle()),
          filterButton(
              context: context,
              text: "Cancelled",
              width: 110,
              onPressed: () => setState(() {
                    selectedIndex = 3;
                  }),
              style: selectedIndex == 3 ? selectedStyle() : unselectedStyle())
        ],
      ),
      Padding(
        padding: const EdgeInsets.all(16),
        child: (() {
          if (selectedIndex == 1) {
            return activeScreen();
          } else if (selectedIndex == 2) {
            return completedScreen();
          } else if (selectedIndex == 3) {
            return cancelledScreen();
          } else {
            return Container(); // Default case
          }
        })(),
      )
    ]));
  }

  Widget foodElem(
      {required String resName,
      required double summaryPrice,
      required int sumaryQuantity,
      required String time}) {
    return SizedBox(
      height: 120,
      width: double.infinity,
      child: Row(
        children: [
          ClipRRect(
            borderRadius: BorderRadius.circular(16),
            child: Image.network(
              "https://fastly.picsum.photos/id/715/200/300.jpg?hmac=jMgGkNrRGTz5pgw27YMTCyozftm33Rw2fPKQU2FypW4",
            ),
          ),
          const SizedBox(
            width: 20,
          ),
          Expanded(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(
                      resName,
                      style: TextStyle(
                          color: Theme.of(context).primaryColor,
                          fontWeight: FontWeight.bold,
                          fontSize: 20),
                    ),
                    Text(
                      "\$${summaryPrice.toStringAsFixed(2)}",
                      style: TextStyle(
                          color: Theme.of(context).colorScheme.secondary,
                          fontWeight: FontWeight.bold,
                          fontSize: 20),
                    ),
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(time,
                        style: TextStyle(
                          color: Theme.of(context).primaryColor,
                        )),
                    Text("$sumaryQuantity items",
                        style: TextStyle(
                          color: Theme.of(context).primaryColor,
                        ))
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: [
                    filterButton(
                        context: context,
                        text: "View",
                        width: 60,
                        onPressed: () {},
                        style: selectedStyle()),
                  ],
                )
              ],
            ),
          )
        ],
      ),
    );
  }

  Widget activeScreen() {
    return Column(
      children: [
        customDivider(context),
        foodElem(
            resName: "Pizza Huh",
            sumaryQuantity: 2,
            summaryPrice: 20.00,
            time: "29 Nov, 01.20 pm"),
        customDivider(context),
        foodElem(
            resName: "MALA",
            sumaryQuantity: 3,
            summaryPrice: 12.80,
            time: "30 Nov, 11.20 pm"),
        customDivider(context),
        foodElem(
            resName: "Chef off",
            sumaryQuantity: 10,
            summaryPrice: 99.99,
            time: "29 Nov, 01.20 pm"),
        customDivider(context),
      ],
    );
  }

  Widget completedScreen() {
    return noHistory(status: "complete");
  }

  Widget cancelledScreen() {
    return Column(
      children: [
        customDivider(context),
        foodElem(
            resName: "7-11",
            sumaryQuantity: 12,
            summaryPrice: 17.8,
            time: "29 Nov, 01.20 pm"),
        customDivider(context),
        foodElem(
            resName: "Dairy King",
            sumaryQuantity: 3,
            summaryPrice: 8.00,
            time: "30 Nov, 11.20 pm"),
        customDivider(context),
      ],
    );
  }

  Widget noHistory({required String status}) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        const SizedBox(
          height: 100,
        ),
        Icon(
          Icons.receipt_long_outlined,
          color: Color(0xFFFFDECF),
          size: 180,
        ),
        Text(
          "You don't have any\n $status orders at this\ntime",
          textAlign: TextAlign.center,
          style: TextStyle(
              color: Theme.of(context).colorScheme.secondary,
              fontSize: 30,
              fontWeight: FontWeight.w500),
        )
      ],
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
