import 'package:get_it/get_it.dart';

import 'app.service.dart';

final getIt = GetIt.instance;

Future<void>  setupServiceLocator() async {
  getIt.registerSingleton<AppModel>(AppModelImplementation());
}