import React from 'react';
import { SafeAreaView, StyleSheet, Text } from 'react-native';
import { FlatList } from 'react-native-gesture-handler';
import ListItem from '../components/ListItem';
import { connect } from 'react-redux';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
});

const ClipScreen = props => {
  return (
    <SafeAreaView style={styles.container}>
      <FlatList
        data={props.user.clips}
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

const mapStateToProps = state => {
  return {
    user: state.user,
  };
};

export default connect(mapStateToProps, null)(ClipScreen);
