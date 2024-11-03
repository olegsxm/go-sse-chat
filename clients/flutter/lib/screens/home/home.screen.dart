import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          backgroundColor: Colors.indigo,
          title: const Text(
            'ChatApp',
            style: TextStyle(color: Colors.white),
          ),
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: () {},
          backgroundColor: Colors.indigo,
          child: const Icon(
            Icons.messenger_outline,
            color: Colors.white,
          ),
        ),
        body: Chats());
  }
}

class Chats extends StatelessWidget {
  Chats();

  List<Dialog> dialogs = [
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
    const Dialog(),
  ];

  @override
  Widget build(BuildContext context) {
    return ListView(
      children: dialogs,
    );
  }
}

class Dialog extends StatelessWidget {
  const Dialog({super.key});

  @override
  Widget build(BuildContext context) {
    return const Row(
      children: [
        SizedBox(
            width: double.infinity,
            child: Text(
              "Message Sender Name",
              textAlign: TextAlign.left,
            )),
        SizedBox(
            width: double.infinity,
            child: Text("Message text with supper helpful information about nothing")),
      ],
    );
  }
}
