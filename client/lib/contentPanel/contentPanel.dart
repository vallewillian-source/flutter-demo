// ignore_for_file: library_private_types_in_public_api, file_names

import 'package:flutter/material.dart';

class ContentPanel extends StatefulWidget {
  const ContentPanel({Key? key}) : super(key: key);

  @override
  _ContentPanelState createState() => _ContentPanelState();
}

class _ContentPanelState extends State<ContentPanel> {
  @override
  Widget build(BuildContext context) {
    return Container(
        height: MediaQuery.of(context).size.height,
        color: const Color(0xff0e0e0e),
        child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.start,
            children: const [Text('.. ..'), Text('.. ..')]));
  }
}
