import { goUp } from "server/goUp";

import { switchSort } from "server/switchSort";

import { directoryPath } from "server/directoryPathSetter";

import { sortFlag } from "server/sortFlagSetter";

import { get } from "server/get";

import "../styles/index.css"

// Вызов функции при завершении загрузки страницы
document.addEventListener('DOMContentLoaded', () => get(directoryPath, sortFlag), false);

// Добавление Слушателей событий
document.getElementById("directoryUp").addEventListener("click", goUp)

document.getElementById("sort").addEventListener("click", switchSort)