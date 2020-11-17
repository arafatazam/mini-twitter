<template>
  <div class="box">
    <article class="media">
      <div class="media-content">
        <div class="content">
          <p>
            <strong>John Smith</strong> <small>@johnsmith</small>
            <small>31m</small>
            <br />
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean
            efficitur sit amet massa fringilla egestas. Nullam condimentum
            luctus turpis.
          </p>
        </div>
        <nav class="level is-mobile">
          <div class="level-left">
            <a class="level-item" aria-label="reply">
              <span class="icon is-small">
                <i class="fas fa-reply" aria-hidden="true"></i>
                <small>230</small>
              </span>
            </a>
          </div>
        </nav>
      </div>
    </article>
  </div>
</template>

<style scoped>
span{
    margin-left: 14px;
}
span>small{
    margin-left: 5px;
    margin-top: 3px;
    color: gray;
}
</style>