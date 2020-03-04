import React from 'react';
import { createAppContainer } from 'react-navigation';
import { createStackNavigator } from 'react-navigation-stack';
import { createBottomTabNavigator } from 'react-navigation-tabs';
import { FontAwesome } from '@expo/vector-icons';
import HomeScreen from '../screens/HomeScreen';
import ArticleScreen from '../screens/ArticleScreen';
import ClipScreen from '../screens/ClipScreen';

const HomeStack = createStackNavigator({
  Home: {
    screen: HomeScreen,
    navigationOptions: {
      title: 'ArticleList',
      headerShown: false,
    },
  },
  Articles: {
    screen: ArticleScreen,
  },
});

const ClipStack = createStackNavigator({
  Clip: {
    screen: ClipScreen,
    navigationOptions: {
      title: 'ClipList',
    },
  },
  Articles: {
    screen: ArticleScreen,
  },
});

const AppNavigator = createBottomTabNavigator({
  Home: {
    screen: HomeStack,
    navigationOptions: {
      tabBarIcon: ({ tintColor }) => {
        return <FontAwesome name="home" color={tintColor} size={24} />;
      },
    },
  },
  Clip: {
    screen: ClipStack,
    navigationOptions: {
      tabBarIcon: ({ tintColor }) => {
        return <FontAwesome name="bookmark" color={tintColor} size={24} />;
      },
    },
  },
});

export default createAppContainer(AppNavigator);
