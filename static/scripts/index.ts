import { goUp } from "./goUp";

import { switchSort } from "./switchSort";

import { directoryPath } from "./directoryPathSetter";

import { sortFlag } from "./sortFlagSetter";

import { get } from "./get";

import "../styles/index.css"

// Вызов функции при завершении загрузки страницы
document.addEventListener('DOMContentLoaded', () => get(directoryPath, sortFlag), false);

// Добавление Слушателей событий
document.getElementById("directoryUp").addEventListener("click", goUp)

document.getElementById("sort").addEventListener("click", switchSort)