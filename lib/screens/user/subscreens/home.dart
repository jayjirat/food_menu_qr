import 'package:flutter/material.dart';
import 'package:flutter/gestures.dart';
import 'package:food_menu_qr/components/custom_divider.dart';
import 'package:food_menu_qr/components/user_stack.dart';

final List<Map<String, dynamic>> partners = [
  {
    "name": "KFD",
    "image_url":
        "https://fastly.picsum.photos/id/425/200/200.jpg?hmac=rC9sY_-TCJnYO9XF-5_pnNdcesi3TZCoWRWhlwSNxcw"
  },
  {
    "name": "McGonald",
    "image_url":
        "https://fastly.picsum.photos/id/242/200/200.jpg?hmac=Z3aa8zbEQkEMFgnVh0Pn96vmCZHhJ17qzCrePYksrcY"
  },
  {
    "name": "PizzaHuh",
    "image_url":
        "https://fastly.picsum.photos/id/188/200/200.jpg?hmac=TipFoTVq-8WOmIswCmTNEcphuYngcdkCBi4YR7Hv6Cw"
  },
  {
    "name": "BurgerQueen",
    "image_url":
        "https://fastly.picsum.photos/id/77/200/200.jpg?hmac=RaFJkrixMn3dR7INSPWcmjC7HCxmggmF5mTlMpyEHsQ"
  }
];

Widget gridBuilderItem(int index) => ClipRRect(
      borderRadius: BorderRadius.circular(16),
      child: GridTile(
        footer: GridTileBar(
          backgroundColor: Colors.black38,
          title: Text(partners[index]['name']),
        ),
        child: Image.network("${partners[index]['image_url']}"),
      ),
    );

Widget home(BuildContext context) {
  return userStack(
    child: ListView(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text("Our Partner Restaurants",
                style: TextStyle(
                    color: Theme.of(context).primaryColor,
                    fontWeight: FontWeight.bold,
                    fontSize: 20)),
            RichText(
              text: TextSpan(
                  text: "View All >",
                  style: TextStyle(
                      color: Theme.of(context).colorScheme.secondary,
                      fontWeight: FontWeight.bold),
                  recognizer: TapGestureRecognizer()..onTap = () => {}),
            )
          ],
        ),
        customDivider(context),
        const SizedBox(
          height: 5,
        ),
        SizedBox(
          height: 150,
          child: GridView.builder(
            shrinkWrap: true,
            scrollDirection: Axis.horizontal,
            itemCount: partners.length,
            gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: 1,
              childAspectRatio: 1.0,
              mainAxisSpacing: 10,
            ),
            itemBuilder: (context, index) {
              return gridBuilderItem(index);
            },
          ),
        ),
        const SizedBox(
          height: 20,
        ),
        SizedBox(
          height: 150,
          child: Container(
            decoration: BoxDecoration(
                color: Theme.of(context).colorScheme.secondary,
                borderRadius: BorderRadius.circular(20)),
            child: Row(
              children: [
                const SizedBox(
                  width: 20,
                ),
                Expanded(
                  child: RichText(
                      text: TextSpan(
                          text: "Experience our\ndelicious new dish\n\n",
                          children: [
                        TextSpan(
                            text: "30% OFF",
                            style: TextStyle(
                                fontSize: 28, fontWeight: FontWeight.bold))
                      ])),
                ),
                Image.network(
                    "https://fastly.picsum.photos/id/107/300/200.jpg?hmac=1ChAQunGwvU9ZitmS9YO3D5NBW_bFoYG1mg_5qWfAZ4")
              ],
            ),
          ),
        ),
        const SizedBox(
          height: 20,
        ),
        Text("Steps to Enjoy",
            style: TextStyle(
                color: Theme.of(context).primaryColor,
                fontWeight: FontWeight.bold,
                fontSize: 20)),
        customDivider(context),
        Padding(
          padding: const EdgeInsets.all(8.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment
                .spaceBetween, // Align the whole row in the center
            children: [
              Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  CircleAvatar(
                    radius: 30,
                    backgroundColor: Theme.of(context)
                        .colorScheme
                        .primary, // Change background color of circle
                    child: Text(
                      "1",
                      style: TextStyle(
                        color: Colors.white,
                        fontWeight: FontWeight.bold,
                        fontSize: 20,
                      ),
                    ),
                  ),
                  SizedBox(height: 8),
                  Text(
                    "Scan",
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                      color: Colors.black,
                    ),
                  ),
                ],
              ),
              Column(
                children: [
                  Icon(
                    Icons.arrow_forward_outlined,
                    color: Theme.of(context)
                        .colorScheme
                        .primary, // Adjust arrow color
                    size: 30,
                  ),
                  const SizedBox(
                    height: 30,
                  )
                ],
              ),
              Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  CircleAvatar(
                    radius: 30,
                    backgroundColor: Theme.of(context).colorScheme.primary,
                    child: Text(
                      "2",
                      style: TextStyle(
                        color: Colors.white,
                        fontWeight: FontWeight.bold,
                        fontSize: 20,
                      ),
                    ),
                  ),
                  SizedBox(height: 8),
                  Text(
                    "Order",
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                      color: Colors.black,
                    ),
                  ),
                ],
              ),
              Column(
                children: [
                  Icon(
                    Icons.arrow_forward_outlined,
                    color: Theme.of(context).colorScheme.primary,
                    size: 30,
                  ),
                  const SizedBox(
                    height: 30,
                  )
                ],
              ),
              Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  CircleAvatar(
                    radius: 30,
                    backgroundColor: Theme.of(context).colorScheme.primary,
                    child: Text(
                      "3",
                      style: TextStyle(
                        color: Colors.white,
                        fontWeight: FontWeight.bold,
                        fontSize: 20,
                      ),
                    ),
                  ),
                  SizedBox(height: 8),
                  Text(
                    "Relax",
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                      color: Colors.black,
                    ),
                  ),
                ],
              ),
            ],
          ),
        )
      ],
    ),
  );
}
