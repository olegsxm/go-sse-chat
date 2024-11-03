import 'package:client/services/app.service.dart';
import 'package:client/services/service-locator.dart';
import 'package:go_router/go_router.dart';

import '../screens/auth/login.screen.dart';
import '../screens/auth/register.screen.dart';
import '../screens/home/home.screen.dart';

final router = GoRouter(
  initialLocation: '/',
  debugLogDiagnostics: true,
  routes: [
    GoRoute(
        name: "home",
        path: "/",
        builder: (context, state) => const HomeScreen()),
    GoRoute(
        name: "login",
        path: "/login",
        builder: (context, state) => LoginScreen()),
    GoRoute(
        name: "sign-up",
        path: "/sign-up",
        builder: (context, state) => const RegisterScreen())
  ],
  redirect: ((context, state) {
    var url = state.uri.toString();
    var userIsEmpty = getIt<AppModel>().user == null;
    var isAuthPages = url.contains("/login") || url.contains("/sign-");

    print("is auth");
    print(isAuthPages);

    if (userIsEmpty && isAuthPages) {
      return null;
    }

    if (!userIsEmpty && isAuthPages) {
      return "/";
    }

    return userIsEmpty ? "/login" : null;
  }),
);
