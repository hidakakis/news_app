import React, { useState, useEffect } from 'react';
// import { StyleSheet, View, FlatList, SafeAreaView } from 'react-native';
import { StyleSheet, View, FlatList } from 'react-native';
import SafeAreaView from 'react-native-safe-area-view';
import ListItem from '../components/ListItem.js';
import Constants from 'expo-constants';
import axios from 'axios';

// const URL = `https://newsapi.org/v2/top-headlines?country=jp&category=business&apiKey=${Constants.manifest.extra.newsApiKey}`;
const URL = 'http://0.0.0.0:5000/api/fetch_latest_articles';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  itemContener: {
    height: 100,
    width: '100%',
    borderColor: 'gray',
    borderWidth: 1,
    flexDirection: 'row',
  },
  leftContener: {
    width: 100,
  },
  rightContener: {
    flex: 1,
    flexDirection: 'column',
    padding: 10,
    justifyContent: 'space-evenly',
  },
  text: {
    fontSize: 16,
  },
  subText: {
    fontSize: 12,
    color: 'gray',
  },
});

// 7eb67555c5e04db195876c61520d5974
export default HomeScreen = props => {
  const [articles, setArticles] = useState([]);
  useEffect(() => {
    fetchArticles();
  }, []);

  const fetchArticles = async () => {
    try {
      const response = await axios.get(URL);
      setArticles(response.data.articles);
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <SafeAreaView style={styles.container}>
      <FlatList
        data={articles}
        renderItem={({ item }) => (
          <ListItem
            author={item.author}
            title={item.title}
            imageUrl={item.urlToImage}
            onPress={() =>
              props.navigation.navigate('Articles', { article: item })
            }
          />
        )}
        keyExtractor={(item, index) => index.toString()}
      />
    </SafeAreaView>
  );
};
