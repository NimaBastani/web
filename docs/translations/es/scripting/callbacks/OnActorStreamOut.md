---
título: OnActorStreamOut
descripción: Este callback se llama cuando un actor se deja de transmitir por el cliente de un jugador.
tags: []
---

<VersionWarn name='callback' version='SA-MP 0.3.7' />

## Descripción

Este callback se llama cuando un actor se deja de transmitir por el cliente de un jugador.

| Nombre      | Descripción                                                    |
| ----------- | -------------------------------------------------------------- |
| actorid     | El ID del actor que dejó de ser transmitido por el jugador.    |
| forplayerid | El ID del jugador que dejó de transmitir al actor.             |

## Devoluciones

Siempre se llama primero en filterscripts.

## Ejemplos

```c
public OnActorStreamOut(actorid, forplayerid)
{
    new string[40];
    format(string, sizeof(string), "El actor %d dejó de ser transmitido a tu jugador.", actorid);
    SendClientMessage(forplayerid, 0xFFFFFFFF, string);
    return 1;
}
```

## Notas

<TipNPCCallbacks />

## Funciones Relacionadas
