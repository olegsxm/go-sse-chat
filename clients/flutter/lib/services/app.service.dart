import 'dart:convert';

import 'package:client/models/user.dart';
import 'package:flutter/material.dart';
import 'package:localstorage/localstorage.dart';

import 'http.dart';


abstract class AppModel extends ChangeNotifier {
  Future<void> signIn (String login, String password);
  User? get user;
  Future<void> singUp(String login, String password);
}

class AppModelImplementation extends AppModel {
  AppModelImplementation();

  @override
  User? get user{
    var u = localStorage.getItem("auth");
    if (u != null) {
      return User.fromJSON(json.decode(u));
    }
    return null;
  }

  @override
  Future<void> signIn(String login, String password) async {;
    final res = await dioI.post("/auth/sign-in", data: {'login': login, 'password': password});
    localStorage.setItem("auth", json.encode(res.data));
  }

  @override
  Future<void> singUp(String login, String password) async {
    final res = await dioI.post("/auth/sign-up", data: {'login': login, 'password': password});
    localStorage.setItem("auth", json.encode(res.data));
  }

}
