# Direct Message Server

이 프로젝트는 고 언어로 작성된 간단한 DM 서비스 Stub입니다.  
자세한 구현을 작성할 지는 잘 모르겠습니다.

## Components

이 프로젝트는 크게 7가지 부분으로 구성됩니다.

1. DirectMessageServer
   - 전송될 메시지를 받는 서버입니다.
   - 어느정도 필요한 인증을 수행합니다.
   - 기존 메시지에 대한 검색을 수행합니다.
2. MessageBroker
   - Kafka나 NATS Jetstream같은 메시지 브로커입니다.
   - 전송 요청된 메시지를 전파하는 역할을 합니다.
3. PushServer
   - 전송 요청된 메시지를 대상 사용자에게 전송합니다.
4. PushMessageChecker
   - 이미 사용자에게 전송된 메시지인지 체크하는 역할을 하는 분산 캐시 혹은 저장소입니다.
5. PushTokenStore
   - 등록된 기기의 토큰 정보를 기록합니다.
   - 전송 시에 조회하는 역할도 동일하게 수행합니다.
6. MessageStorage
   - 메시지를 저장하는 곳입니다.
   - cassandra, scylladb, S3, clickhouse, mongo, elasticsearch 등을 활용합니다.
7. MessageIndexer
   - 메시지를 검색하는 컴포넌트입니다.
-    elasticsearch, meilisearch, postgres 등을 활용합니다
