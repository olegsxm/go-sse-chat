class User {
  String id;
  String login;
  String token;

  User({required this.id, required this.login, required this.token});

  factory User.fromJSON(Map<String, dynamic> json) => User(
      id: json["user"]["id"],
      login: json["user"]["login"],
      token: json["token"]
  );
}
