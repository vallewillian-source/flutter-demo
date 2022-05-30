// ignore_for_file: library_private_types_in_public_api, prefer_const_literals_to_create_immutables

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:sofalab/contentPanel/contentPanel.dart';
import 'package:sofalab/sidePanel/sidePanel.dart';
import 'package:sofalab/widgets/split.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  SystemChrome.setPreferredOrientations([
    DeviceOrientation.portraitUp,
    DeviceOrientation.portraitDown,
  ]).then((value) => runApp(const MyApp()));
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var baseTheme = ThemeData(
        brightness: Brightness.dark,
        fontFamily: "Roboto",
        textTheme: const TextTheme().copyWith(
          bodyText1: const TextStyle(fontSize: 12, fontFamily: 'Roboto'),
          bodyText2: const TextStyle(fontSize: 12, fontFamily: 'Roboto'),
        ));

    return MaterialApp(theme: baseTheme, home: const HomePage());
  }
}

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final _sidePanelScrollController = ScrollController();
  final _contentPanelcrollController = ScrollController();

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Container(
            height: MediaQuery.of(context).size.height,
            color: const Color(0xff191818),
            child: Split(
                axis: Axis.horizontal,
                initialFirstFraction: 0.2,
                firstChild: Scrollbar(
                  controller: _sidePanelScrollController,
                  child: SingleChildScrollView(
                      controller: _sidePanelScrollController,
                      child: const SidePanel()),
                ),
                secondChild: Scrollbar(
                  controller: _contentPanelcrollController,
                  child: SingleChildScrollView(
                      controller: _contentPanelcrollController,
                      child: const ContentPanel()),
                ))));
  }
}
