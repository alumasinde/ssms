<?php

namespace Core;

use DateTime;

class Session
{
    public static function start(): void
    {
        if (session_status() === PHP_SESSION_NONE) {
            ini_set('session.cookie_httponly', '1');
            ini_set('session.cookie_samesite', 'Lax');
            session_name('sms_sess');
            session_start();
        }
    }
    public static function can(string $permission): bool
{
    $permissions = self::get('permissions', []);
    return in_array($permission, $permissions);
}

public static function hasRole()

    public static function set(string $key, mixed $value): void
    {
        self::start();
        $_SESSION[$key] = $value;
    }

    public static function get(string $key, mixed $default = null): mixed
    {
        self::start();
        return $_SESSION[$key] ?? $default;
    }

    public static function has(string $key): bool
    {
        self::start();
        return isset($_SESSION[$key]);
    }

    public static function remove(string $key): void
    {
        self::start();
        unset($_SESSION[$key]);
    }

    public static function destroy(): void
    {
        self::start();
        session_unset();
        session_destroy();
    }

    public static function isLoggedIn(): bool
    {
        return self::has('access_token') && self::has('user');
    }

    public static function user(): ?array
    {
        return self::get('user');
    }

    public static function flash(string $key, mixed $value = null): mixed
    {
        if ($value !== null) {
            self::set('_flash_' . $key, $value);
            return null;
        }
        $val = self::get('_flash_' . $key);
        self::remove('_flash_' . $key);
        return $val;
    }

    public static function formatDate(?string $date, string $format = 'Y-m-d'): string
{
    if (empty($date)) {
        return '';
    }

    return (new DateTime($date))->format($format);
}
}
