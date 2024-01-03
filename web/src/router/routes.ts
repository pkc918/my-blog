import { RouteRecordRaw } from "vue-router";
const Layout = () => import("@/components/Layout.vue");
const Intro = () => import("@/views/Intro.vue");
const Blogs = () => import("@/views/Blogs.vue");
const Projects = () => import("@/views/Projects.vue");
const Photo = () => import("@/views/Photo.vue");
const Music = () => import("@/views/Music.vue");

const routes: RouteRecordRaw[] = [
  {
    path: "/:catchAll(.*)",
    name: "404",
    meta: {
      title: "404 - 建设中"
    },
    component: Layout,
  },
  {
    path: "/",
    name: "root",
    component: Layout,
    redirect: { name: "intro" },
    children: [
      {
        path: "/",
        name: "intro",
        meta: {
          title: "陪我去看海吧"
        },
        component: Intro,
      },
      {
        path: "/blog",
        name: "blog",
        meta: {
          title: "Blog - 陪我去看海吧"
        },
        component: Blogs,
      },
      {
        path: "/projects",
        name: "projects",
        meta: {
          title: "Projects - 陪我去看海吧"
        },
        component: Projects,
      },
      {
        path: "/photo",
        name: "photo",
        meta: {
          title: "Photo - 陪我去看海吧"
        },
        component: Photo,
      },
      {
        path: "/music",
        name: "music",
        meta: {
          title: "Music - 陪我去看海吧"
        },
        component: Music,
      },
    ],
  },
];

export default routes;
