<?php

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Session;

class AuthController extends Controller
{
    public function login(Request $request)
    {
        // Validate incoming request (email and password)
        $request->validate([
            'email' => 'required|email',
            'password' => 'required',
        ]);

        // Send login credentials to third-party API for authentication
        $response = Http::post(env('API_URL').'/login', [
            'email' => $request->email,
            'password' => $request->password,
        ]);
        $data = $response->object();
        // Check if the third-party API returned a successful response
        if ($data->message != "Success"){
        return redirect()->route('login')
            ->withErrors(['gagal' => $data->message])
            ->withInput();
        }else{
            session()->regenerate();
            $minutes = 720;  // 12 hours in minutes
            session(['expires_at' => now()->addMinutes($minutes)]); // Store expiration time in session
            $user = $data->data;
            // Optionally, you can store more data in the session
            session(['name' => $user->name]);
            session(['position' => $user->position]);

            return redirect()->intended('/beranda');
        }
    }

    public function loginView(Request $request)
    {
        if(session('name')) {
            return redirect()->intended('/beranda');
        }
        return view('login');
    }

    public function logout(Request $request)
    {
        // Clear all session data (destroy the session completely)
        session()->flush();

        // Regenerate the session ID to prevent session fixation attacks
        session()->regenerate();

        // Optionally, you can redirect the user to the login page with a success message
        return redirect()->route('login')->with('status', 'You have been logged out successfully.');
    }
    

    public function beranda(Request $request)
    {
        if(!session('name')) {
            return redirect()->intended('/login');
        }
        return view('beranda');
    }
}
