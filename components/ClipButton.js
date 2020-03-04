import React from 'react';
import { StyleSheet, TouchableOpacity } from 'react-native';
import { FontAwesome } from '@expo/vector-icons';

const styles = StyleSheet.create({
  container: {
    padding: 5,
  },
});

export default ClipButton = ({ onPress, enabled }) => {
  const name = enabled ? 'bookmark' : 'bookmark-o';
  return (
    <TouchableOpacity onPress={onPress} style={styles.container}>
      <FontAwesome name={name} size={40} collor="gray" />
    </TouchableOpacity>
  );
};
