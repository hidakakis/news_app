import React from 'react';
import { StyleSheet, Text, View, Image, TouchableOpacity } from 'react-native';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
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

const ListItem = ({ imageUrl, title, author, onPress }) => {
  return (
    <TouchableOpacity style={styles.itemContener} onPress={onPress}>
      <View style={styles.leftContener}>
        <Image
          style={{ width: 100, height: 100 }}
          source={
            imageUrl != null
              ? { uri: imageUrl }
              : { uri: '../image/no_image_react.png' }
          }
        />
      </View>
      <View style={styles.rightContener}>
        <Text numberOfLines={3} style={styles.text}>
          {title}
        </Text>
        <Text style={styles.subText}>{author}</Text>
      </View>
    </TouchableOpacity>
  );
};

export default ListItem;
