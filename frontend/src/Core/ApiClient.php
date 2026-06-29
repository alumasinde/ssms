<?php

namespace Core;

class ApiClient
{
    private string $baseUrl;
    private ?string $token;

    public function __construct(?string $token = null)
    {
        $this->baseUrl = Config::apiUrl();
        $this->token   = $token ?? Session::get('access_token');
    }

    public function get(string $path, array $query = []): array
    {
        $url = $this->baseUrl . $path;
        if ($query) $url .= '?' . http_build_query(array_filter($query, fn($v) => $v !== null && $v !== ''));
        return $this->request('GET', $url);
    }

    public function post(string $path, array $data = []): array
    {
        return $this->request('POST', $this->baseUrl . $path, $data);
    }

    public function put(string $path, array $data = []): array
    {
        return $this->request('PUT', $this->baseUrl . $path, $data);
    }

    public function patch(string $path, array $data = []): array
    {
        return $this->request('PATCH', $this->baseUrl . $path, $data);
    }

    public function delete(string $path): array
    {
        return $this->request('DELETE', $this->baseUrl . $path);
    }

    private function request(string $method, string $url, array $data = []): array
    {
        $headers = [
            'Content-Type: application/json',
            'Accept: application/json',
            'Host: ' . Config::currentDomain(),
        ];

        if ($this->token) {
            $headers[] = 'Authorization: Bearer ' . $this->token;
        }

        $ch = curl_init($url);
        curl_setopt_array($ch, [
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_HTTPHEADER     => $headers,
            CURLOPT_CUSTOMREQUEST  => $method,
            CURLOPT_TIMEOUT        => 15,
            CURLOPT_CONNECTTIMEOUT => 5,
        ]);

        if (in_array($method, ['POST', 'PUT', 'PATCH']) && $data) {
            curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
        }

        $response = curl_exec($ch);
        $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        $error    = curl_error($ch);
        curl_close($ch);

        if ($error) {
            return ['success' => false, 'error' => 'API connection failed: ' . $error, 'data' => null];
        }

        $decoded = json_decode($response, true);

        if ($httpCode === 401) {
            Session::destroy();
            header('Location: /auth/login');
            exit;
        }

        if ($httpCode === 403) {
            return ['success' => false, 'error' => 'You do not have permission to perform this action.', 'data' => null];
        }

        if ($httpCode === 404) {
            return ['success' => false, 'error' => 'Resource not found.', 'data' => null];
        }

        if ($httpCode >= 500) {
            return ['success' => false, 'error' => 'Server error. Please try again later.', 'data' => null];
        }

        return $decoded ?? ['success' => false, 'error' => 'Invalid API response (HTTP ' . $httpCode . ')', 'data' => null];
    }
}
