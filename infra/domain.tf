resource "aws_acm_certificate" "ecs_domain_certificate" {
  domain_name       = "*.${var.ecs_domain_name}"
  validation_method = "DNS"

  tags = {
    Name = "${var.ecs_cluster_name}-Certificate"
  }
}

# data：aws_route53（パブリックゾーン）は手動で既に用意してあるので、既にあるものを参照する形式
data "aws_route53_zone" "ecs_domain" {
  name         = var.ecs_domain_name # 探したいドメイン名
  private_zone = false # パブリックゾーンを使用、ECSサービスを外部公開するため  # 検索条件：パブリックゾーンを探す
}

# Note:DNSレコードを使った検証 => SSL通信で使用する SSL証明書を使うには「このドメインの持ち主が本当にあなたですか？」という検証が必要
# X-server の時のように以下の対応を行なっている
# 例: example.com というドメインのSSL証明書を取得する場合

# ① AWSのACM（証明書サービス）が検証用の情報を発行します：
#   - レコード名: _acme-challenge.example.com
#   - レコードの値: abc123def456（AWSが生成した特別な文字列）
#   - レコードタイプ: CNAME

# ② この情報をDNSレコードとしてRoute 53に登録します：
#   _acme-challenge.example.com → abc123def456

# ③ AWSが「確かにDNSレコードが登録されている」と確認
#   →「このドメインの持ち主に間違いない」と検証完了
resource "aws_route53_record" "cert_validation" {
  for_each = {
    for ecs in aws_acm_certificate.ecs_domain_certificate.domain_validation_options : ecs.domain_name => {
      name   = ecs.resource_record_name
      record = ecs.resource_record_value
      type   = ecs.resource_record_type
    }
  }

  # Note: allow_overwrite = true
  # for_eachで作成したマップは、eachとして参照できます

  # 同じ名前のDNSレコードが既に存在する場合の動作を制御します
  # trueに設定すると：

  # 既存のDNSレコードがあっても、新しい値で上書きします
  # 証明書の更新時に必要な新しい検証レコードを自動的に作成できます


  # もしfalseの場合：

  # 既存レコードがあるとエラーになります
  # 証明書の更新が失敗する可能性があります
  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60 # TTL（Time To Live）: DNSレコードがキャッシュされる時間
  type            = each.value.type
  zone_id         = data.aws_route53_zone.ecs_domain.zone_id
}

resource "aws_acm_certificate_validation" "ecs_domain_certificate_validation" {
  certificate_arn         = aws_acm_certificate.ecs_domain_certificate.arn
  validation_record_fqdns = [for record in aws_route53_record.cert_validation : record.fqdn]


}
