<template>  
    <div>  
      <h1>{{ message }}</h1>  
      <button @click="changeMessage">Change Message</button>  
    </div>  
  </template>  
  
  <script>  
  import { useMainStore } from '../store';  
  
  export default {  
    setup() {  
      const store = useMainStore();  
  
      const changeMessage = () => {  
        store.updateMessage('Message updated!');  
      };  
  
      return {  
        message: store.message,  
        changeMessage,  
      };  
    },  
  };  
  </script>  