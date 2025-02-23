import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class Home extends ConsumerStatefulWidget {
  const Home({super.key});

  @override
  HomeState createState() => HomeState();
}

class HomeState extends ConsumerState<Home> {
  int selectedIndex = 0;
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

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        toolbarHeight: 150,
        automaticallyImplyLeading: false,
        title: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                children: [
                  Expanded(
                    child: TextField(
                      decoration: InputDecoration(
                          hintText: "Search...",
                          prefixIcon: Icon(Icons.search),
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(30),
                            borderSide: BorderSide.none,
                          ),
                          filled: true,
                          fillColor: Colors.white,
                          contentPadding: const EdgeInsets.all(0)),
                    ),
                  ),
                  const SizedBox(
                    width: 20,
                  ),
                  InkWell(
                    onTap: () {},
                    child: Container(
                      height: 30,
                      width: 30,
                      decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(20),
                      ),
                      child: Center(
                        child: Icon(
                          Icons.notifications_outlined,
                          color: Theme.of(context).colorScheme.secondary,
                          size: 25,
                        ),
                      ),
                    ),
                  ),
                  const SizedBox(
                    width: 10,
                  ),
                  InkWell(
                    onTap: () {},
                    child: Container(
                      height: 30,
                      width: 30,
                      decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(20),
                      ),
                      child: Center(
                        child: Icon(
                          Icons.person_outline,
                          color: Theme.of(context).colorScheme.secondary,
                          size: 25,
                        ),
                      ),
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 12),
              RichText(
                text: TextSpan(
                  text: "Welcome, ",
                  style: TextStyle(
                      fontSize: 26,
                      fontWeight: FontWeight.bold,
                      color: Colors.white),
                  children: [
                    TextSpan(
                      text: "Fullname\n",
                      style: TextStyle(
                          fontWeight: FontWeight.bold,
                          color: Theme.of(context).colorScheme.secondary),
                    ),
                    TextSpan(
                      text:
                          "Rise and shine! It's time to enjoy something delicious.",
                      style: TextStyle(
                          fontSize: 12,
                          color: Theme.of(context).colorScheme.secondary),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
      body: Center(
        child: Stack(
          children: [
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
                      topLeft: Radius.circular(30),
                      topRight: Radius.circular(30)),
                  color: Colors.white),
              child: Padding(
                padding: const EdgeInsets.all(16),
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
                                  color:
                                      Theme.of(context).colorScheme.secondary,
                                  fontWeight: FontWeight.bold),
                              recognizer: TapGestureRecognizer()
                                ..onTap = () => {}),
                        )
                      ],
                    ),
                    Divider(
                      color: Theme.of(context).colorScheme.primary,
                    ),
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
                                      text:
                                          "Experience our\ndelicious new dish\n\n",
                                      children: [
                                    TextSpan(
                                        text: "30% OFF",
                                        style: TextStyle(
                                            fontSize: 28,
                                            fontWeight: FontWeight.bold))
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
                    Divider(
                      color: Theme.of(context).colorScheme.primary,
                    ),
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
                                backgroundColor:
                                    Theme.of(context).colorScheme.primary,
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
                                backgroundColor:
                                    Theme.of(context).colorScheme.primary,
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
              ),
            )
          ],
        ),
      ),
      bottomNavigationBar: ClipRRect(
        borderRadius: BorderRadius.only(
            topLeft: Radius.circular(40), topRight: Radius.circular(40)),
        child: BottomNavigationBar(
          type: BottomNavigationBarType.fixed,
          backgroundColor: Theme.of(context).colorScheme.secondary,
          unselectedItemColor: Colors.white,
          selectedItemColor: Theme.of(context).colorScheme.primary,
          currentIndex: selectedIndex,
          onTap: (index) {
            setState(() {
              selectedIndex = index;
            });
          },
          items: [
            BottomNavigationBarItem(
              icon: Icon(Icons.home_outlined),
              label: "Home",
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.history_outlined),
              label: "History",
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.qr_code_outlined),
              label: "QR code",
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.favorite_outline),
              label: "Favorites",
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.support_agent_outlined),
              label: "Support",
            ),
          ],
        ),
      ),
    );
  }

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
}
