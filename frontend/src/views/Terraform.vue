<template>  
    <div>  
      <h1>Terraform Management</h1>  
      <button @click="initTerraform">Initialize</button>  
      <button @click="applyTerraform">Apply</button>  
      <button @click="destroyTerraform">Destroy</button>  
    </div>  
  </template>  
  
  <script>  
  import axios from 'axios'  
  
  export default {  
    methods: {  
      async initTerraform() {  
        await axios.post('/api/terraform/init')  
        alert('Terraform initialized')  
      },  
      async applyTerraform() {  
        await axios.post('/api/terraform/apply')  
        alert('Terraform applied')  
      },  
      async destroyTerraform() {  
        await axios.post('/api/terraform/destroy')  
        alert('Terraform destroyed')  
      },  
    },  
  }  
  </script>  