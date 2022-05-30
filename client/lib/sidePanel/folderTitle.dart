// ignore_for_file: library_private_types_in_public_api, file_names, prefer_const_literals_to_create_immutables

import 'package:flutter/material.dart';

class SidePanelFolderTitle extends StatefulWidget {
  final String title;
  const SidePanelFolderTitle({Key? key, required this.title}) : super(key: key);

  @override
  _SidePanelFolderTitleState createState() => _SidePanelFolderTitleState();
}

PageController page = PageController();

class _SidePanelFolderTitleState extends State<SidePanelFolderTitle> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return const Text(
      "PinkApp",
      style: TextStyle(
        fontSize: 12,
        fontWeight: FontWeight.bold,
        color: Color.fromARGB(255, 179, 179, 179),
      ),
    );
  }
}
