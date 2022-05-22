<template>
  <div id="app">
    <h1>To-Do List</h1>
    <div id="widget-container"></div>
    <form @submit.prevent="sendItem" class="todo-form">
      <input
        id="todo-input"
        class="todo-input"
        type="text"
        size="50"
        v-model="todoitem"
        placeholder="Add a new item"
      />
      <button id="submit" type="submit" value="Submit">Submit</button>
    </form>
    <ul class="card">
      <div>
      <template v-for="(item, index) in todolist" :key="index">
          <div class="todo-list-content">
            <span class="one-todo-list-content-title">
              {{ index+1+":" }}{{ item }}
            </span>
            <span class = "one-todo-list-content">
                <span id="edit" v-on:click = "updateItem(index)">Update</span>
            </span>
            <span class = "one-todo-list-content">
                <span id="danger" v-on:click = "deleteItem(index)">Delete</span>
            </span>
          </div>
      </template>
      </div>
    </ul>
    <div>{{ message }}</div>
  </div>
</template>

<script>
import axios from "axios";

const appData = {
  todolist: ["More to do"],
  todoitem: "",
  message: ""
};

const baseUrl = "http://localhost:6060";

export default {
  name: "App",
  data() {
    return appData;
  },
  methods: {
    getList: getList,
    sendItem: sendItem,
    deleteItem: deleteItem,
    updateItem: updateItem,
  },
  mounted: function() {
    this.getList();
  },
  watch: {
    '$appData.todolist': {
      handler() {
        this.getList();
      },
      immediate: true
    } 
  }
}

function deleteItem(index) {
  console.log(this.todolist)
  console.log(appData.todolist)
  axios.delete(`${baseUrl}/api/lists/${index}`)
    .then(function() { 
      getList();
    })
    .catch(function (error) {
      appData.todolist = [error.message];
    });
}

function updateItem(index) {
  const params = new URLSearchParams();
  const val = this.todoitem;
  params.append("item", val);
  axios.put(`${baseUrl}/api/lists/${index}`, params)
   .then(function () {
      appData.todolist[index] = val;
      getList();
   })
   .catch();
}

function getList() {
  axios.get(`${baseUrl}/api/lists`).then(res => {
    appData.todolist = res.data.list;
    const inputVal = document.getElementById("todo-input");
    inputVal.value = "";
  });
}

async function sendItem() {
  const params = new URLSearchParams();
  const val = this.todoitem;
  params.append("item", val);
  await axios.post(`${baseUrl}/api/lists`, params).then(
    function () {
      appData.todolist.push(val);
    }
  ).catch(function (error) {
    appData.todolist = [error.message];
  });

}


</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.todo-form{
  margin: 0 auto;
  box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
  transition: 0.3s;
  border-radius: 5px;
  padding: 5px;
  width: 400px; 
}

.todo-input{
  height: 24px;
  margin-left: 6px;
  margin-right: 6px;
  width: 320px;
  outline: none;
  /* border: none; */
}

#submit {
  height: 30px;
  width: 60px;
  border: none;
  outline: none;
  background-color: #eee;
  border-radius: 5px;
}

#submit:hover {
  background-color: #ccc;
  cursor: pointer;
}

#danger {
  display: inline-block;
  height: 24px;
  line-height: 24px;
  width: 60px;
  border: none;
  outline: none;
  background-color: red;
  color: white;
  border-radius: 5px;
}

#danger:hover {
  background-color: lightcoral;
  cursor: pointer;
}

#edit {
  display: inline-block;
  height: 24px;
  line-height: 24px;
  width: 60px;
  border: none;
  outline: none;
  background-color: blue;
  color: white;
  border-radius: 5px;
}

#edit:hover {
  background-color: skyblue;
  cursor: pointer;
}
.card {
  margin: 0 auto;
  box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
  transition: 0.3s;
  border-radius: 5px;
  padding: 5px;
  width: 400px;
}

.todo-list-content{
  margin: 5px;
}

.one-todo-list-content{
  margin: 5px;
}

.one-todo-list-content-title{
  display: inline-block;
  margin: 5px;
  width: 240px;
  text-align: left;
}

</style>
