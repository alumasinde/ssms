<?php

namespace Core;

class Config
{
    private static array $data = [];
    private static bool $loaded = false;

    public static function load(string $envFile = null): void
    {
        if (self::$loaded) return;

        $file = $envFile ?? dirname(__DIR__) . '/.env';
        if (file_exists($file)) {
            $lines = file($file, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
            foreach ($lines as $line) {
                if (str_starts_with(trim($line), '#')) continue;
                [$key, $value] = array_pad(explode('=', $line, 2), 2, '');
                self::$data[trim($key)] = trim($value);
            }
        }
        self::$loaded = true;
    }

    public static function get(string $key, mixed $default = null): mixed
    {
        self::load();
        return self::$data[$key] ?? $_ENV[$key] ?? getenv($key) ?: $default;
    }

    public static function apiUrl(): string
    {
        return rtrim(self::get('APP_URL', 'http://localhost:8080/api/v1'), '/');
    }

    public static function appName(): string
    {
        return self::get('APP_NAME', 'SchoolMS');
    }

    /**
     * The domain the browser is currently on, without port.
     * e.g. ssms.highway.localhost:3000 → ssms.highway.localhost
     * This gets forwarded to Go as the Host header so tenant resolution works.
     */
    public static function currentDomain(): string
    {
        $host = $_SERVER['HTTP_HOST'] ?? '';
        // Strip port number
        return preg_replace('/:\d+$/', '', $host);
    }
}