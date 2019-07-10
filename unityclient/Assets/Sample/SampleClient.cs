using System;
using Grpc.Core;
using Sample.Chat;
using UnityEngine;

public class SampleClient : MonoBehaviour
{
    // Start is called before the first frame update
    async void Start()
    {
        var channel = new Channel("localhost", 10000, ChannelCredentials.Insecure);
        var client = new RouteGuide.RouteGuideClient(channel);

        var replay = await client.HelloAsync(new Name {Name_ = "takashi"}).ResponseAsync;
        
        Debug.Log(replay.Message);
    }
}
