<?php
/*
 * wycms -  A PHP Framework For Web
 *
 * @package  Wycms
 * @author   superl <superl@nepenthes.cn>
 */
define('SUPCMS_START',microtime(true));

// 调试级别
define('SUPCMS_ENVIRONMENT', 'development');

// web根目录
define('SUPCMS_PATH', dirname(__FILE__). DIRECTORY_SEPARATOR);

/*
|--------------------------------------------------------------------------
| 引导启动
|--------------------------------------------------------------------------
*/
$app = require_once __DIR__ . '/system/nepenthes.php';