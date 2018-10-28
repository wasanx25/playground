<template>
  <div>
    <div id="header">
      <div class="container">
        <h1>wasanx25 playground</h1>
      </div>
    </div>
    <div id="main" class="container">
      <nuxt/>
    </div>
    <div id="footer" class="container">
      <p>Copyright © 2018 wasanx25 All Right Reserved.</p>
    </div>
  </div>
</template>

<style>
body {
  background-color: #4db6ac;
}

#header {
  background-color: #212121;
}

#header .container {
  padding: 12px 0;
}

#header h1 {
  margin: 0;
  font-size: 1.8rem;
  color: #efefef;
}

#main {
  height: 80vh;
}

#footer {
  text-align: center;
}

#footer p {
  color: #9e9e9e;
}
</style>
