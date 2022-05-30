// ignore_for_file: library_private_types_in_public_api, file_names, prefer_const_literals_to_create_immutables

import 'package:flutter/material.dart';

class SidePanelActionTitle extends StatefulWidget {
  final String title;
  const SidePanelActionTitle({Key? key, required this.title}) : super(key: key);

  @override
  _SidePanelActionTitleState createState() => _SidePanelActionTitleState();
}

PageController page = PageController();

class _SidePanelActionTitleState extends State<SidePanelActionTitle> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return MouseRegion(
        cursor: SystemMouseCursors.click,
        child: Text(
          widget.title,
          style: const TextStyle(
              fontSize: 12, color: Color.fromARGB(255, 179, 179, 179)),
        ));
  }
}
