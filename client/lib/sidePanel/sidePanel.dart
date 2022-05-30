// ignore_for_file: library_private_types_in_public_api, file_names, prefer_const_literals_to_create_immutables

import 'package:flutter/material.dart';
import 'package:sofalab/sidePanel/actionTitle.dart';
import 'package:sofalab/sidePanel/folderTitle.dart';
import 'package:sofalab/sidePanel/title.dart';
import 'package:sofalab/widgets/split.dart';
import 'package:sofalab/widgets/tree-view/flutter_simple_treeview.dart';
import 'package:sofalab/widgets/tree-view/src/primitives/tree_node.dart';

class SidePanel extends StatefulWidget {
  const SidePanel({Key? key}) : super(key: key);

  @override
  _SidePanelState createState() => _SidePanelState();
}

PageController page = PageController();

class _SidePanelState extends State<SidePanel> {
  final _endpointsScrollController = ScrollController();
  final _resultsScrollController = ScrollController();

  final TextEditingController _textController = TextEditingController(text: '''
{
  "PinkApp": [
    {
      "title": "send a message"
    },
    {
      "title": "create a channel"
    }
  ]
}
''');

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
        height: MediaQuery.of(context).size.height,
        color: const Color(0xff191818),
        child: Split(
            axis: Axis.vertical,
            initialFirstFraction: 0.6,
            firstChild: Scrollbar(
              controller: _endpointsScrollController,
              child: SingleChildScrollView(
                  controller: _endpointsScrollController,
                  child: Padding(
                      padding:
                          const EdgeInsets.only(left: 15, top: 15, right: 5),
                      child: Column(children: [
                        Row(children: [
                          Expanded(
                              child: OutlinedButton(
                                  onPressed: () {},
                                  child: const Text("import actions")))
                        ]),
                        const SidePanelTitle(title: 'ACTIONS', count: "2"),
                        TreeView(
                          nodes: [
                            TreeNode(
                              content:
                                  const SidePanelFolderTitle(title: "PinkApp"),
                              children: [
                                TreeNode(
                                    content: const SidePanelActionTitle(
                                  title: "send a message",
                                )),
                                TreeNode(
                                    content: const SidePanelActionTitle(
                                  title: "create a channel",
                                ))
                              ],
                            ),
                          ],
                          indent: 0,
                        )
                      ]))),
            ),
            secondChild: Scrollbar(
                controller: _resultsScrollController,
                child: SingleChildScrollView(
                  controller: _resultsScrollController,
                  child: Column(children: [
                    const Padding(
                        padding: EdgeInsets.only(
                            left: 15, right: 5, top: 5, bottom: 0),
                        child: Divider(
                          height: 1,
                          thickness: 1,
                          indent: 0,
                          endIndent: 0,
                          color: Color.fromARGB(255, 54, 54, 54),
                        )),
                    Padding(
                        padding:
                            const EdgeInsets.only(left: 15, top: 0, right: 5),
                        child: Column(children: [
                          Row(children: [
                            const SidePanelTitle(title: 'RESULTS', count: "0")
                          ])
                        ]))
                  ]),
                ))));
  }
}
