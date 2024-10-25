import 'package:dio/dio.dart';

final options = BaseOptions(
  baseUrl: 'http://10.0.2.2:8080/api',
  responseType: ResponseType.json,
);

final dioI = Dio(options);