package com.example.chat

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.Stable
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import androidx.navigation.NavGraphBuilder
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import androidx.navigation.navigation
import com.example.chat.ui.theme.ChatTheme

@Composable
fun rememberAppState(
    scaffoldState: ScaffoldState = rememberScaffoldState(),
    navController: NavHostController = rememberNavController()
) =
    remember(scaffoldState, navController) {
        AppState(scaffoldState, navController)
    }

@Stable
class AppState(
    val scaffoldState: ScaffoldState,
    val navController: NavHostController
) {}

object MainDestinations {
    const val HOME_ROUTE = "home"
    const val GAME_CARD_DETAIL_ROUTE = "cardRoute"
    const val GAME_CARD = "gameCard"
    const val SUB_CATALOG_ROUTE = "subCatalog"
    const val CATALOG_GAME = "catalogGame"
}

fun NavGraphBuilder.addHomeGraph(
    modifier: Modifier = Modifier
) {
//    composable(HomeSections.CATALOG.route) {
//        CatalogScreen()
//    }
//    composable(HomeSections.PROFILE.route) {
//        ProfileScreen()
//    }
//    composable(HomeSections.SEARCH.route) {
//        SearchScreen()
//    }
}

fun NavGraphBuilder.navGraph() {
    navigation(
        route = MainDestinations.HOME_ROUTE,
        startDestination = ""
    ) {
        addHomeGraph()
    }
}

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        enableEdgeToEdge()
        setContent {
            ChatTheme {
                val navController = rememberNavController()
                Scaffold( modifier = Modifier.fillMaxSize() )
                {
                    innerPadding -> Greeting( name = "Android", modifier = Modifier.padding(innerPadding)),
                    NavHost(navController = navController, startDestination = "/") {
                        newGraph()
                    }
                }
            }
        }
    }
}

@Composable
fun Greeting(name: String, modifier: Modifier = Modifier) {
    Text(
        text = "Hello $name!",
        modifier = modifier
    )
}

@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    ChatTheme {
        Greeting("Android")
    }
}