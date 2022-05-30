// ignore_for_file: library_private_types_in_public_api, file_names, prefer_const_literals_to_create_immutables

import 'package:flutter/material.dart';

class SidePanelTitle extends StatefulWidget {
  final String title;
  final String count;
  const SidePanelTitle({Key? key, required this.title, required this.count})
      : super(key: key);

  @override
  _SidePanelTitleState createState() => _SidePanelTitleState();
}

PageController page = PageController();

class _SidePanelTitleState extends State<SidePanelTitle> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: const EdgeInsets.only(top: 30, bottom: 10),
        child: Row(children: [
          Text(
            widget.title,
            style: const TextStyle(fontSize: 11, fontWeight: FontWeight.bold),
          ),
          const SizedBox(width: 10),
          Chip(
            padding: const EdgeInsets.all(0.0),
            label: Text(
              widget.count,
              style: const TextStyle(fontSize: 10, fontWeight: FontWeight.w200),
            ),
          )
        ]));
  }
}
